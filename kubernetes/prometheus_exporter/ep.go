package main

import (
	//"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	//"os"
	"regexp"
	"strings"
	"time"
	"github.com/cactus/go-statsd-client/statsd"
	s "gitlab-odx.oracledx.com/sauron/endpoints-checker/pkg/statsd"
	//stat "github.com/smira/go-statsd"
	"github.com/Jeffail/gabs"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"sync"
	"sort"
	"net/url"
	"strconv"
	"crypto/tls"
)

type EndpointConfig struct {
	EndpointRequestInterval         int      `yaml:"EndpointRequestInterval"`
	UptimeRobotScrapIntervalMinutes int      `yaml:"UptimeRobotScrapIntervalMinutes"`
	ExcludeEndPoints                []string `yaml:"ExcludeEndpoints"`
}

type Monitor struct {
	Url          string
	FriendlyName string
	HttpUsername string
	HttpPassword string
	FailureCount int
	SuccessCount int
	Https        bool
	MetricName   string
	NewMonitor   bool
}

type logWriter struct {
}

func (m Monitor) String() string {
	return fmt.Sprintf("[%s, %s]", m.Url, m.FriendlyName)
}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(time.Now().UTC().Format("[2006-01-02T15:04:05.999Z]") + " [INFO] " + string(bytes))
}

func init() {
	log.SetOutput(new(logWriter))
}

const (
	ConfigFile = "/etc/endpoints-checker/excluded_endpoints.txt"
)

func (c *EndpointConfig) getEndpointConfig() *EndpointConfig {
	yamlFile, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}

func main() {
	monitorMap = map[string]*Monitor{}
	eliminatedMonitorMap = map[string]string{}
	duplicateMonitors = map[string]string{}
	channles = map[string]chan (bool){}

	s.StatsdExporter()

	config = EndpointConfig{
		EndpointRequestInterval:         60,
		UptimeRobotScrapIntervalMinutes: 10,
	}

	if _, err := os.Stat(ConfigFile); os.IsNotExist(err) {
		y, err := yaml.Marshal(config)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(string(y))

		err1 := ioutil.WriteFile(ConfigFile, y, 0644)
		if err1 != nil {
			log.Printf("Config post error  #%v ", err)
		}
	}

	go StartServer()

	statsdClient, err := statsd.NewClient("127.0.0.1:9125", "endpoints-checker")
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	if err != nil {
		panic(err.Error())
	}
	defer statsdClient.Close()

	trInsecure := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClientInsecure := &http.Client{
		Transport: trInsecure,
		Timeout:   time.Second * 60,
	}
	httpClientSecure := &http.Client{
		Timeout: time.Second * 60,
	}

	getAllMonitors(httpClientInsecure, httpClientSecure, statsdClient)
}

var re = regexp.MustCompile("https?://(.*?)\\.(prod|dev|integ)\\.(.*?)\\.sauron\\.(.*?)(/.*|$)")
var reMeta = regexp.MustCompile("https?://(.*?)\\.(prod|dev|integ)\\.sauron\\.(.*?)(/.*|$)")
var goRoutineCount int
var mutex sync.Mutex

func checkEndpoint(chStop <-chan bool, statsdClient statsd.Statter, httpClientSecure *http.Client, httpClientInsecure *http.Client, monitor *Monitor, config EndpointConfig) { //, excludedEndpoints map[string]bool) {
	mutex.Lock()
	goRoutineCount++
	mutex.Unlock()
	for {
		select {
		case x := <-chStop:
			mutex.Lock()
			goRoutineCount--
			mutex.Unlock()
			log.Println("Ending thread:", monitor.Url, x)
			return
		default:
		}

		var httpErr error
		req, _ := http.NewRequest(http.MethodGet, monitor.Url, nil)
		if monitor.HttpUsername != "" && monitor.HttpPassword != "" {
			req.SetBasicAuth(monitor.HttpUsername, monitor.HttpPassword)
		}

		var resp *http.Response
		if strings.HasPrefix(monitor.Url, "https://") {
			resp, httpErr = httpClientSecure.Do(req)
		} else {
			resp, httpErr = httpClientInsecure.Do(req)
		}

		for count := 0; count < 3 && (((resp == nil) || (resp != nil && resp.StatusCode != 200)) || httpErr != nil); count++ {
			log.Printf("\tRetrying the request [count:%d, URL: %s, httpError: %s]\n", count, monitor.Url, httpErr)
			if strings.HasPrefix(monitor.Url, "https://") {
				resp, httpErr = httpClientSecure.Do(req)
			} else {
				resp, httpErr = httpClientInsecure.Do(req)
			}
			if resp != nil && resp.StatusCode == 200 {
				log.Printf("\t####Retry success:[count:%d, status code: %d, URL: %s]\n", count, resp.StatusCode, monitor.Url)
				monitor.SuccessCount ++
				monitor.NewMonitor = false
				break
			}
			monitor.FailureCount ++
			time.Sleep(time.Duration((count+1)*5) * time.Second)
		}

		errMsg := ""
		if httpErr != nil {
			errMsg = fmt.Sprintf("Error loading %s: %v\n", monitor.Url, httpErr)
		} else if resp.StatusCode != 200 {
			errMsg = fmt.Sprintf("Unexpected response code for %s. Expected %d, Actual %d\n", monitor.Url, 200, resp.StatusCode)
		}

		if monitor.NewMonitor == false {
			if errMsg != "" {
				log.Printf("--------------------Querying URL failed: %v-----------------------, Error = %v\n", monitor.Url, errMsg)
				//monitor.FailureCount ++
				statsdClient.Raw("health."+monitor.MetricName, "0|g", 1.0)
			} else {
				//log.Printf("Querying URL succeeded: %v\n", url)
				statsdClient.Raw("health."+monitor.MetricName, "1|g", 1.0)
				monitor.SuccessCount ++
			}
		}
		time.Sleep(time.Duration(config.EndpointRequestInterval) * time.Second)
	}
}

//UPTIME_MONITOR_API_KEY=u492353-8d09a11654a19a89bb343fdb

var monitorMap map[string]*Monitor
var eliminatedMonitorMap map[string]string
var duplicateMonitors map[string]string
var channles map[string]chan (bool)
var config EndpointConfig

func getAllMonitors(httpClientInsecure, httpClientSecure *http.Client, statsdClient statsd.Statter) {
	for {
		log.Printf("Refreshing data from uptime monitor\n")
		config.getEndpointConfig()
		allMonitors := getMonitorsFromFile()
		//allMonitors := getMonitorsFromUptime()
		var newMonitors []*Monitor
		var eliminatedMonitors []*Monitor
		totalEndpoints := 0
		for _, monitor := range allMonitors {
			totalEndpoints ++
			url := monitor.Url
			if strings.Contains(url, "help.") || strings.Contains(url, "test") {
				eliminatedMonitors = append(eliminatedMonitors, monitor)
				continue
			}

			https := false
			if strings.HasPrefix(url, "https://") {
				https = true
			}

			sauronType := "classic"
			if strings.HasPrefix(monitor.FriendlyName, "http") || strings.HasPrefix(monitor.FriendlyName, "canary") {
				sauronType = "operator"
			}

			var metricName string
			//if config.EnableUrlFilter == true {
			matches := re.FindStringSubmatch(url)
			if len(matches) == 0 {
				// Special case for Meta Sauron environments
				matches = reMeta.FindStringSubmatch(url)
				if len(matches) > 3 {
					strArr := strings.Split(matches[1], ".")
					metricNameSub := strings.Replace(strings.Replace(matches[3], "-", "_", -1), ".", "_", -1)
					metricName = ""
					if len(strArr) == 2 {
						metricName = fmt.Sprintf("%s.%s.%s.%s", sauronType, matches[2], matches[2]+"_"+strArr[1]+"_sauron_"+metricNameSub, strArr[0])
					} else {
						metricName = fmt.Sprintf("%s.%s.%s.%s", sauronType, matches[2], matches[2]+"_sauron_"+metricNameSub, strArr[0])
					}
				} else {
					eliminatedMonitors = append(eliminatedMonitors, &Monitor{url, monitor.FriendlyName, "", "",
						0, 0, false, "", false})
					continue
				}
			} else if len(matches) > 4 {
				strArr := strings.Split(matches[1], ".")
				metricNameSub := strings.Replace(strings.Replace(matches[4], "-", "_", -1), ".", "_", -1)
				envName := strings.Replace(matches[3], "-", "_", -1)
				metricName = ""
				if len(strArr) == 2 {
					metricName = fmt.Sprintf("%s.%s.%s.%s", sauronType, matches[2], matches[2]+"_"+envName+"_"+strArr[1]+"_sauron_"+metricNameSub, strArr[0])
				} else {
					metricName = fmt.Sprintf("%s.%s.%s.%s", sauronType, matches[2], matches[2]+"_"+envName+"_sauron_"+metricNameSub, strArr[0])
				}
			} else {
				eliminatedMonitors = append(eliminatedMonitors, &Monitor{url, monitor.FriendlyName, "", "",
					0, 0, false, "", false})
				continue
			}
			//	}

			failureCount := 0
			successCount := 0
			if _, ok := monitorMap[url]; ok {
				failureCount = monitorMap[url].FailureCount
				successCount = monitorMap[url].SuccessCount
			}

			monitor.SuccessCount = successCount
			monitor.FailureCount = failureCount
			monitor.Https = https
			monitor.MetricName = metricName
			monitor.NewMonitor = false
			newMonitors = append(newMonitors, monitor)
		}

		//-------------------
		// Remove excepted endpoints
	loop:
		for i := 0; i < len(newMonitors); i++ {
			url := newMonitors[i].Url
			for _, rem := range config.ExcludeEndPoints {
				if url == rem {
					newMonitors = append(newMonitors[:i], newMonitors[i+1:]...)
					i-- // Important: decrease index
					continue loop
				}
			}
		}

		//Invoke go routines for first time
		if len(monitorMap) == 0 {
			log.Println("initializing monitors for the firsttime:", len(newMonitors))
			for _, monitor := range newMonitors {
				if _, ok := monitorMap[monitor.Url]; ok {
					log.Println("Found a duplicate url from uptimetobot", monitor.Url)
					duplicateMonitors[monitor.Url] = monitor.FriendlyName
					continue
				}
				monitorMap[monitor.Url] = monitor
				ch := make(chan bool)
				channles[monitor.Url] = ch
				go checkEndpoint(ch, statsdClient, httpClientSecure, httpClientInsecure, monitor, config)
			}
			continue
		}

		//Eliminated Endpoints
		for _, em := range eliminatedMonitors {
			str := strings.Replace(strings.Replace(em.Url, ".", "_", -1), "-", "_", -1)
			strArray := strings.Split(str, "/")
			eliminatedMonitorMap[em.Url] = strArray[2]
		}

		var oldMonitors []*Monitor
		for _, val := range monitorMap {
			oldMonitors = append(oldMonitors, val)
		}
		monitorsToRemove := findDifferenceInMonitors(oldMonitors, newMonitors)
		monitorsToAdd := findDifferenceInMonitors(newMonitors, oldMonitors)

		log.Println("Monitors to add:", monitorsToAdd)
		for _, monitor := range monitorsToAdd {
			monitor.NewMonitor = true
			monitorMap[monitor.Url] = monitor
			log.Println(monitor.Url)
			ch := make(chan bool)
			channles[monitor.Url] = ch
			go checkEndpoint(ch, statsdClient, httpClientSecure, httpClientInsecure, monitor, config)
		}

		log.Println("Monitors to remove:", monitorsToRemove)
		for _, monitor := range monitorsToRemove {
			delete(monitorMap, monitor.Url)
			ch := channles[monitor.Url]
			go func() {
				ch <- true
			}()
			delete(channles, monitor.Url)
		}
		log.Printf("***************************************************\n")
		log.Printf("EndpointRequestInterval:%d secs", config.EndpointRequestInterval)
		log.Printf("UptimeRobotScrapIntervalMinutes:%d mins", config.UptimeRobotScrapIntervalMinutes)
		log.Printf("Total endpoints received from UptimeRobot:%d", totalEndpoints)
		log.Printf("Current Go-routine Count: %d\n", goRoutineCount)
		log.Printf("Current monitors %v\n", len(newMonitors))
		log.Printf("Endpoints, which are not monitored: %v\n", len(eliminatedMonitors))
		log.Printf("Current monitorsMap %v\n", len(monitorMap))
		log.Printf("Current channels running%v\n", len(channles))
		log.Printf("Duplicate Endpoints received from UptimeRobot: %d\n", len(duplicateMonitors))
		log.Printf("***************************************************\n")
		time.Sleep(time.Duration(config.UptimeRobotScrapIntervalMinutes) * time.Minute)
	}
}

//This function is for testing
func getMonitorsFromFile() []*Monitor {
	body, err := ioutil.ReadFile("output.json")
	if err != nil {
		log.Printf("Json .Get err   #%v ", err)
	}
	jsonParsed, _ := gabs.ParseJSON(body)
	values, err := jsonParsed.Children()
	var monitors []*Monitor
	totalEndpoints := 0
	for _, value := range values {
		totalEndpoints ++
		url := value.Path("url").Data().(string)
		username := value.Path("http_username").Data().(string)
		password := value.Path("http_password").Data().(string)
		friendlyName := value.Path("friendly_name").Data().(string)

		monitors = append(monitors, &Monitor{url, friendlyName, username, password,
			0, 0, false, "", false})
	}
	return monitors
}

func getMonitorsFromUptime() []*Monitor {
	offset := 0
	limit := 50
	myurl := "https://api.uptimerobot.com/v2/getMonitors"
	totalEndpoints := 0
	var monitors []*Monitor
	str := "["
	for {
		resp, err := http.PostForm(myurl, url.Values{"api_key": {os.Getenv("UPTIME_MONITOR_API_KEY")},
			"ssl": {"true"}, "format": {"json"}, "logs": {"0"}, "offset": {strconv.Itoa(offset)}, "limit": {strconv.Itoa(limit)}})
		if err != nil {
			continue
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)

		jsonParsed, _ := gabs.ParseJSON(body)
		pagination_limit := int(jsonParsed.Path("pagination.limit").Data().(float64))
		pagination_total := int(jsonParsed.Path("pagination.total").Data().(float64))
		pagination_offset := int(jsonParsed.Path("pagination.offset").Data().(float64))
		if pagination_total > len(monitors) {
			offset = pagination_offset + pagination_limit
			values, _ := jsonParsed.Path("monitors").Children()
			for _, value := range values {
				str = fmt.Sprintf("%s %v ,\n", str, value)
				totalEndpoints ++
				url := value.Path("url").Data().(string)
				username := value.Path("http_username").Data().(string)
				password := value.Path("http_password").Data().(string)
				friendlyName := value.Path("friendly_name").Data().(string)

				monitors = append(monitors, &Monitor{url, friendlyName, username, password,
					0, 0, false, "", false})
			}
		}

		if pagination_total == len(monitors) {
			break
		}

		if offset+limit > pagination_total {
			offset = offset % pagination_total
			limit = pagination_total - len(monitors)
		}
	}

	fmt.Println("Madhu:", monitors)
	return monitors
}

func findDifferenceInMonitors(a, b []*Monitor) (differenceMoniror []*Monitor) {
	log.Println("Verifying for remove monitors.", len(a), len(b))
	l1 := len(a)
	l2 := len(b)
	for i := 0; i < l1; i++ {
		var j int
		for j = 0; j < l2; j++ {
			if (a[i].Url == b[j].Url) {
				break;
			}
		}

		if (j == l2) {
			differenceMoniror = append(differenceMoniror, a[i])
		}
	}
	return differenceMoniror
}

func StartServer() {
	r := gin.Default()
	r.Use(Cors())

	v1 := r.Group("/")
	{
		v1.GET("/health", GetHealth)
		v1.GET("/status", GetEndPoints)
		v1.GET("/config", GetConfig)
		v1.POST("/config", PostConfig)
		v1.GET("/eliminatedmonitors", GetEliminatedEndPoints)
		v1.GET("/duplicatemonitors", GetDuplicateEndPoints)
	}
	log.Println("Starting REST server")
	r.Run(":8080")
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func GetHealth(c *gin.Context) {
	// Display JSON result
	c.JSON(200, "OK")
}

func GetEndPoints(c *gin.Context) {
	monitors := []*Monitor{}
	for _, val := range monitorMap {
		monitors = append(monitors, val)
	}

	sortByErrorCount(monitors)
	c.JSON(200, monitors)
}

func GetConfig(c *gin.Context) {
	c.YAML(200, config.getEndpointConfig())
}

func PostConfig(c *gin.Context) {
	x, _ := ioutil.ReadAll(c.Request.Body)
	log.Printf("PostConfig %s \n", x)
	var config1 *EndpointConfig
	err := yaml.Unmarshal(x, &config1)
	if err != nil {
		log.Printf("Config.Post err   #%v ", err)
	}

	err1 := ioutil.WriteFile(ConfigFile, x, os.ModePerm)
	if err1 != nil {
		log.Printf("Config post error  #%v ", err)
	}

	config = *config1
}

func GetDuplicateEndPoints(c *gin.Context) {
	c.JSON(200, duplicateMonitors)
}
func GetEliminatedEndPoints(c *gin.Context) {
	c.JSON(200, eliminatedMonitorMap)
}

func sortByErrorCount(monitor []*Monitor) {
	sort.SliceStable(monitor, func(i, j int) bool {
		mi, mj := monitor[i], monitor[j]
		switch {
		case mi.FailureCount != mj.FailureCount:
			return mi.FailureCount > mj.FailureCount
		default:
			return mi.FailureCount > mj.FailureCount
		}
	})
}
