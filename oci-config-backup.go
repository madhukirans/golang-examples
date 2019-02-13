package main

import (
	"context"
	"fmt"
	//"io"
	"os"
	"archive/tar"
	_ "github.com/oracle/oci-go-sdk/common"
	_ "github.com/oracle/oci-go-sdk/example/helpers"
	_ "github.com/oracle/oci-go-sdk/objectstorage"

	"github.com/oracle/oci-go-sdk/objectstorage"
	//"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/example/helpers"
	//"strings"
	//"bytes"
	//"io"
	"compress/gzip"
	//"bytes"
	"github.com/oracle/oci-go-sdk/common"
	//"io"
	//"io/ioutil"
	"io"
	"bytes"

	//"crypto/rsa"
	"io/ioutil"
	"strings"
)

var ENV_NAME, BACKUP_ID, OPERATION, BUCKET_NAME, NAMESPACE, INDEX string
func main() {
	ENV_NAME = os.Getenv("ENV_NAME")
	BACKUP_ID = os.Getenv("BACKUP_ID")
	OPERATION = os.Getenv("OPERATION")
	BUCKET_NAME = os.Getenv("BUCKET_NAME")
	NAMESPACE = os.Getenv("NAMESPACE")
	INDEX = os.Getenv("INDEX")

	ENV_NAME="sauron-operator"
	BACKUP_ID="20180011-mad12"
	OPERATION="data-restore"
	BUCKET_NAME="backupmadhu"
	NAMESPACE="odx-sre"
	INDEX="0"

	//stream, size :=
		tar1("xxx")

//	os.Setenv("TF_VAR_private_key_path", "ocid1.user.oc1..aaaaaaaapeeg3yjjpyz32c5b7plo4vqzgttqgqnjluntnatchb5b4ulkxieq")
//	os.Setenv("TF_VAR_fingerprint", "11:e7:7b:b5:52:20:97:92:65:8f:28:4a:ea:5b:fb:ea")
//	os.Setenv("TF_VAR_user_ocid", "ocid1.user.oc1..aaaaaaaapeeg3yjjpyz32c5b7plo4vqzgttqgqnjluntnatchb5b4ulkxieq")
//	os.Setenv("TF_VAR_tenancy_ocid", "ocid1.tenancy.oc1..aaaaaaaaqzv7hhoe4sypzyitsoubjf6hbmr26sxrw452p4slslarsbqz25bq")
//	os.Setenv("TF_VAR_region", "eu-frankfurt-1")
//	//os.Setenv("TF_VAR_u", "/Users/madhuseelam/.oraclebmc/bmcs_api_key.pem.madhu")
//
//
//
//	/*
//	user=ocid1.user.oc1..aaaaaaaapeeg3yjjpyz32c5b7plo4vqzgttqgqnjluntnatchb5b4ulkxieq
//fingerprint=11:e7:7b:b5:52:20:97:92:65:8f:28:4a:ea:5b:fb:ea
//key_file=/Users/madhuseelam/.oraclebmc/bmcs_api_key.pem.madhu
//compartment=ocid1.compartment.oc1..aaaaaaaaoukqmttw335fk3fv6ttaw7hvuf7o6d6sofhchdmofcpe34usvkpa
//tenancy=ocid1.tenancy.oc1..aaaaaaaaqzv7hhoe4sypzyitsoubjf6hbmr26sxrw452p4slslarsbqz25bq
//region=eu-frankfurt-1
//	 */
	ociSourceDir :=  os.Getenv("HOME") + "/.oci"
	ociConfigDistDir :=  os.Getenv("HOME") + "/.oraclebmc"
	ociConfigDistConfig := ociConfigDistDir + "/config"

	if _, err := os.Stat(ociConfigDistConfig); os.IsNotExist(err) {
		os.MkdirAll(ociConfigDistDir, os.ModePerm)

		source, err := ioutil.ReadFile(ociSourceDir + "/config")
		if err != nil {
		//	return fmt.Errorf("Could not read oci config file.%s/config ", ociSourceDir )
		}
		configStr := string(source)
		if !strings.Contains(configStr, "key_file") {
			configStr += "\nkey_file=/Users/madhuseelam/madhu/oraclebmc/bmcs_api_key.pem.madhu\n"//\nkey_file=" + ociSourceDir + "/private_key\n"
		}
		fmt.Println(configStr)
		err = ioutil.WriteFile(ociConfigDistConfig, []byte(configStr), 0644)
		if err != nil {
		//	return fmt.Errorf("Could not create oci config file: %s ", ociConfigDistConfig)
		}
	}

	prov := common.DefaultConfigProvider()
	fmt.Println(prov.PrivateRSAKey())
	os.Exit(0)

	storageClinet, clerr := objectstorage.NewObjectStorageClientWithConfigurationProvider(common.DefaultConfigProvider())

	helpers.FatalIfError(clerr)
	ctx := context.Background()
	//objectname := "kiran.tar.gz"
	//
	//request := objectstorage.PutObjectRequest{
	//	NamespaceName: &NAMESPACE,
	//	BucketName:    &BUCKET_NAME,
	//	ObjectName:    &objectname,
	//	ContentLength: &size,
	//	PutObjectBody: stream,
	//	//OpcMeta:       metadata,
	//}
	//_, err := storageClinet.PutObject(ctx, request)
	//fmt.Println("put object", err)
	//namespace := context.getNamespace(ctx, storageClinet)
	//storageClinet.ListObjects()
	////context.
	//
	//
	//fmt.Println( listObjects(ctx, storageClinet, NAMESPACE, BUCKET_NAME, "") )
	//
	fmt.Println( putObject(ctx, storageClinet, NAMESPACE, BUCKET_NAME, "kiran.tar.gz", "/tmp/xxx"))

	//OPERATION= strings.ToLower(OPERATION)
	//if OPERATION  == "full-config-backup" {
	//	fmt.Println("Backing up configs based on ${CONFIG_VOLUME}...")
	//	// For each file in this directly, we take the Sauron name to the be string up the last "-", and the
	//	cd ${CONFIG_VOLUME}
	//	for f in *; do
	//	FILE_NAME=`echo $f | rev | cut -d"-" -f1  | rev`
	//	oci os object put --force --namespace ${NAMESPACE} -bn ${BUCKET_NAME} --name ${ENV_NAME}/config/${BACKUP_ID}/${SAURON_NAME}/${FILE_NAME} --file $f
	//	done
	//}

	//if OPERATION == "data-restore" {
	//	component := os.Getenv("COMPONENT")
	//	COMPONENT_FILE_NAME := component + index
	//	SAURON_NAME:=os.Getenv("SAURON_NAME")
	//	pathPrefix :=  ENV_NAME + "/data/" + SAURON_NAME + "/" + BACKUP_ID + "/"
	//	fmt.Println("Restoring ${COMPONENT}...")
	//
	//	objectList := listObjects(ctx, c, namespace, BUCKET_NAME, "mseelam-so/config/sauron-1/20181016-144546/")
	//	for _, obj := range objectList {
	//		if obj == pathPrefix + COMPONENT_FILE_NAME + ".tar.gz" {
	//
	//		}
	//	}
	//	//# Stream download to tar for extraction
	//	//cd ${DATA_VOLUME}
	//	status=`oci os object list --namespace ${NAMESPACE} -bn ${BUCKET_NAME} --prefix ${ENV_NAME}/data/${SAURON_NAME}/${BACKUP_ID}/${COMPONENT_FILE_NAME}.tar.gz`
	//	if [[ $status =~ "${ENV_NAME}/data/${SAURON_NAME}/${BACKUP_ID}/${COMPONENT_FILE_NAME}.tar.gz" ]]; then
	//	echo 'Backup Found: ${ENV_NAME}/data/${SAURON_NAME}/${BACKUP_ID}/${COMPONENT_FILE_NAME}.tar.gz'
    //    oci os object get --namespace ${NAMESPACE} -bn ${BUCKET_NAME} --name ${ENV_NAME}/data/${SAURON_NAME}/${BACKUP_ID}/${COMPONENT_FILE_NAME}.tar.gz --file - | tar zxv
	//	else
	//	fmt.Println("Backup Not Found: ${ENV_NAME}/data/${SAURON_NAME}/${BACKUP_ID}/${COMPONENT_FILE_NAME}.tar.gz")
	//	fmt.Println("Restoring first backup ${ENV_NAME}/data/${SAURON_NAME}/${BACKUP_ID}/${COMPONENT_FILE_NAME}.tar.gz")
	//	COMPONENT_FILE_NAME=COMPONENT + "-0"
	//	//oci os object get --namespace ${NAMESPACE} -bn ${BUCKET_NAME} --name ${ENV_NAME}/data/${SAURON_NAME}/${BACKUP_ID}/${COMPONENT_FILE_NAME}.tar.gz --file - | tar zxv
	//	fi
	//
	//	# Perform any component-specific fixups
	//	if COMPONENT == "elasticsearch" {
	//		useradd elasticsearch
	//		chown -R elasticsearch ${DATA_VOLUME}
	//	}
	//} else if OPERATION == "data-backup" {
	//	COMPONENT=COMPONENT + INDEX
	//	//SAURON_NAME=SAURON_NAME:?Variable SAURON_NAME must be defined}"
	//	fmt.Println("Backing up ${COMPONENT}...")
	//
	//	# Tar and stream to upload
	//	//cd ${DATA_VOLUME}
	//	//tar zcv . | oci os object put --force --namespace ${NAMESPACE} -bn ${BUCKET_NAME} --name ${ENV_NAME}/data/${SAURON_NAME}/${BACKUP_ID}/${COMPONENT}.tar.gz --file -
	//} else if OPERATION == "config-backup" {
	//	fmt.Println("Backing up configs based on ${CONFIG_VOLUME}...")
	//
	//	# For each file in this directly, we take the Sauron name to the be string up the last "-", and the
	//	//cd ${CONFIG_VOLUME}
	//	/*for f in *; do
	//	  SAURON_NAME=`echo $f | rev | cut -d"-" -f2-  | rev`
	//	  FILE_NAME=`echo $f | rev | cut -d"-" -f1  | rev`
	//	  oci os object put --force --namespace ${NAMESPACE} -bn ${BUCKET_NAME} --name ${ENV_NAME}/config/${SAURON_NAME}/${BACKUP_ID}/${FILE_NAME} --file $f
	//	done
	//	*/
	//}
}

// ExampleObjectStorage_UploadFile shows how to create a bucket and upload a file
func ExampleObjectStorage_UploadFile() {
//	objects = getObject(ctx, c, namespace, BUCKET_NAME, "mseelam-so/config/sauron-1/20181016-144546/")
	//defer deleteBucket(ctx, c, namespace, bname)
	//
	//contentlen := 1024 * 1000
	//filepath, filesize := writeTempFileOfSize(int64(contentlen))
	//filename := path.Base(filepath)
	//defer func() {
	//	os.Remove(filename)
	//}()
	//
	//file, e := os.Open(filepath)
	//defer file.Close()
	//helpers.FatalIfError(e)
	//
	//e = putObject(ctx, c, namespace, bname, filename, filesize, file, nil)
	//helpers.FatalIfError(e)
	//defer deleteObject(ctx, c, namespace, bname, filename)

	// Output:
	// get namespace
	// create bucket
	// put object
	// delete object
	// delete bucket
}

func listObjects(ctx context.Context, c objectstorage.ObjectStorageClient, namespace, bname string, prefix string) ([]string) {
	request1 := objectstorage.ListObjectsRequest{
		BucketName: &bname,
		NamespaceName: &namespace,
		Prefix: &prefix,
	}

	l, r1 := c.ListObjects(ctx, request1)

	var list []string

	for _, obj := range l.Objects {
		list = append(list, *obj.Name)
	}
	fmt.Println(list, " Error: ",r1)
	return list
}

func getNamespace(ctx context.Context, c objectstorage.ObjectStorageClient) string {
	request := objectstorage.GetNamespaceRequest{}
	r, err := c.GetNamespace(ctx, request)
	helpers.FatalIfError(err)
	fmt.Println("get namespace")
	return *r.Value
}

func putObject(ctx context.Context, c objectstorage.ObjectStorageClient, namespace, bucketname, objectname string, filepath string) error { //contentLen int64, content io.ReadCloser, metadata map[string]string) error {
	file, _ := os.Open(filepath)
	defer file.Close()

	i, e := os.Stat(filepath)
	if e != nil {
		return e
	}
	// get the size
	size := i.Size()

	request := objectstorage.PutObjectRequest{
		NamespaceName: &namespace,
		BucketName:    &bucketname,
		ObjectName:    &objectname,
		ContentLength: &size,
		PutObjectBody: file,
		//OpcMeta:       metadata,
	}
	_, err := c.PutObject(ctx, request)
	fmt.Println("put object")
	return err
}

func tar1(filepath string)  (io.ReadCloser, int64) {

	var files = []struct {
		Name, Body string
	}{
		{"a/readme.txt", "This archive contains some text files."},
		{"b/gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"b/todo.txt", "Get animal handling license."},
	}

	var buf bytes.Buffer
	// tar write
	tw := tar.NewWriter( &buf )
	defer tw.Close()

	for _, file := range files {
		hdr := &tar.Header{
			Name: file.Name,
			Mode: 0600,
			Size: int64(len(file.Body)),
		}
		if err := tw.WriteHeader(hdr); err != nil {
			fmt.Println(err)
		}
		if _, err := tw.Write([]byte(file.Body)); err != nil {
			fmt.Println(err)
		}
	}

	// Create and add some files to the archive.
	fw, _ := os.Create( "/tmp/xxx" )
	defer fw.Close()


	// gzip write
	gw := gzip.NewWriter( fw ) //&buf
	defer gw.Close()
	fw.Write(buf.Bytes())



	fr, _ := os.Open( "/tmp/xxx" )
	i, e := os.Stat("/tmp/xxx")
	if e != nil {
		//return e
	}
	// get the size
	size := i.Size()

	return fr, int64(size)

	//if _, err = io.Copy( tw, fr ); err!=nil {
	//	fmt.Println("1",err)
	//}
	//
	//if err := tw.Close(); err != nil {
	//	fmt.Println(err)
	//}
}