package main

import (
	"github.com/oracle/oci-go-sdk/core"
	"github.com/oracle/oci-go-sdk/common"
	"fmt"
	//"github.com/golang/glog"
	//"sort"
	"context"
	//"github.com/oracle/oci-go-sdk/example/helpers"
	"log"
	"sort"
	"github.com/golang/glog"
)

func GetBlockVolumeBackupId(compartmentId string, displayName string) (*string, error) {
	coreClinet, clienrError := core.NewBlockstorageClientWithConfigurationProvider(common.DefaultConfigProvider())
	if clienrError != nil {
		return nil, fmt.Errorf("unable to create OCI clinet: %s", clienrError)
	}
	ctx := context.Background()

	fmt.Println("Getting volume backups for displayName: ", displayName)
	request := core.ListVolumeBackupsRequest{
		CompartmentId: &compartmentId,
		DisplayName:   &displayName,
	}

	// to show how pagination works, reduce number of items to return in a paginated "List" call
	request.Limit = common.Int(1)

	listVolumeBackupFunc := func(request core.ListVolumeBackupsRequest) (core.ListVolumeBackupsResponse, error) {
		return coreClinet.ListVolumeBackups(ctx, request)
	}
	var volBackups []core.VolumeBackup
	for r, err := listVolumeBackupFunc(request); ; r, err = listVolumeBackupFunc(request) {
		if err != nil {
			return nil, fmt.Errorf("unable to get volume backup list error: %v", err)
		}

		log.Printf("list volume backups returns: %v", len(r.Items))

		for _, i := range r.Items {
			volBackups = append(volBackups, i)
		}

		if r.OpcNextPage != nil {
			request.Page = r.OpcNextPage
		} else {
			break
		}
	}

	sort.Slice(volBackups, func(i, j int) bool { return volBackups[i].TimeCreated.Time.Before(volBackups[j].TimeCreated.Time) })
	for _, i := range volBackups {
		glog.Info(*i.Id, *i.TimeCreated)
	}
	if len(volBackups) > 0 {
			return volBackups[len(volBackups)-1].Id, nil
	}
//	fmt.Println(*volBackups[0].DisplayName)

	return nil, fmt.Errorf("block volume backups not found for displayname: %s", displayName)
}

func main() {

	// blockVolumeBackupId, err := pkgsauron.GetBlockVolumeBackupId("
	// ocid1.compartment.oc1..aaaaaaaam4diu23et2xemhpvsbte5lep5ozhfpai6ar6tdgn3723wl7ncxza
	// ocid1.compartment.oc1..aaaaaaaaxwdx6sjt6qqo742lbje7wc3jql6hoby4v7lcheguhrkbs2zcegua
	//        "sauron-ci-741844-backup-prometheus-0-20190410-054401")
	//    fmt.Println("blockVolumeBackupId: ", blockVolumeBackupId)

	fmt.Println(GetBlockVolumeBackupId("ocid1.compartment.oc1..aaaaaaaaxhin4ffkzaml5o7icqbpcxrsgtveczejc3u7d63cr5aigpza6qqq",
	       "sauron-dev-test1-elasticsearch-0-20190410-055107"))

	//fmt.Println(GetBlockVolumeBackupId("ocid1.compartment.oc1..aaaaaaaam4diu23et2xemhpvsbte5lep5ozhfpai6ar6tdgn3723wl7ncxza",
//		"sauron-ci-742380-backup-prometheus-0-20190410-175035")) //sauron-ci-741844-backup-prometheus-0-20190410-054401"))
}
