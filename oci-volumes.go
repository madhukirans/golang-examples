package main

import (
	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/core"

	"context"
	"fmt"
	"github.com/golang/glog"
	"sort"
	//	"time"
)

func main() {

	//vol, _ :=  GetBlockVolumeId("ocid1.compartment.oc1..aaaaaaaaoukqmttw335fk3fv6ttaw7hvuf7o6d6sofhchdmofcpe34usvkpa", "sauron-sauron-1-prometheus-7")
	//createBlockVolumeBackup (vols, "Madhu")
	//fmt.Println(vol)

	//	ListBlockVolumeBackups ("ocid1.compartment.oc1..aaaaaaaaoukqmttw335fk3fv6ttaw7hvuf7o6d6sofhchdmofcpe34usvkpa", "sauron-sauron-1-prometheus-7")
	DeleteBlockVolumeBackups("ocid1.compartment.oc1..aaaaaaaaoukqmttw335fk3fv6ttaw7hvuf7o6d6sofhchdmofcpe34usvkpa")
}

func GetBlockVolumeId(compartmentId string, displayName string) (*core.Volume, error) {
	coreClinet, clienrError := core.NewBlockstorageClientWithConfigurationProvider(common.DefaultConfigProvider())
	if clienrError != nil {
		return nil, fmt.Errorf("Unable to create OCI clinet: %s", clienrError)
	}
	ctx := context.Background()

	glog.Info("Getting volumes for displayName: ", displayName)
	request := core.ListVolumesRequest{
		CompartmentId: &compartmentId,
		DisplayName:   &displayName,
	}
	volList, err := coreClinet.ListVolumes(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("Unable to get List Volumes error: %s", err)
	}

	vols := volList.Items
	glog.Info("Listing volumes for displayName: ", displayName, len(vols))

	sort.Slice(vols, func(i, j int) bool {
		return vols[i].TimeCreated.Time.Before(vols[j].TimeCreated.Time)
	})

	for _, i := range vols {
		fmt.Println(*i.Id, *i.TimeCreated)
	}

	if len(vols) > 0 {
		return &vols[len(vols)-1], nil
	}

	return nil, fmt.Errorf("Block Volumes not found for displayname: %s", displayName)
}

func createBlockVolumeBackup(vol core.Volume, prefix string) error {

	coreClinet, clienrError := core.NewBlockstorageClientWithConfigurationProvider(common.DefaultConfigProvider())
	if clienrError != nil {
		return fmt.Errorf("Unable to create OCI clinet: %s", clienrError)
	}
	ctx := context.Background()

	glog.Info("Creating backup from volume: ", vol.DisplayName, vol.Id)
	//dimsplayName := prefix + "-" + *vol.DisplayName + "-" + *vol.Id
	request := core.CreateVolumeBackupRequest{
		CreateVolumeBackupDetails: core.CreateVolumeBackupDetails{
			VolumeId: vol.Id,
			//DisplayName: &dimsplayName,
		},
	}
	response, err := coreClinet.CreateVolumeBackup(ctx, request)
	if err != nil {
		return fmt.Errorf("Unable to create block volume backup for display name; %s, id: %s, err:%s ", vol.DisplayName, vol.Id, err)
	}
	fmt.Println("Backup initiated for volumes displayName:%s id: %s response: %s", vol.DisplayName, vol.Id, response)

	return nil
}

func ListBlockVolumeBackups(compartmentId string, displayName string) (string, error) {
	coreClinet, clienrError := core.NewBlockstorageClientWithConfigurationProvider(common.DefaultConfigProvider())
	if clienrError != nil {
		return "", fmt.Errorf("Unable to create OCI clinet: %s", clienrError)
	}
	ctx := context.Background()

	glog.Info("List backup volumes for compartment: ", compartmentId, displayName)

	request := core.ListVolumeBackupsRequest{
		CompartmentId: &compartmentId,
		DisplayName:   &displayName,
	}
	response, err := coreClinet.ListVolumeBackups(ctx, request)
	if err != nil {
		return "", fmt.Errorf("Unable to list volume backup for compartment: %s, err:%s ", compartmentId, err)
	}

	fmt.Println("Volume backups for compartmentId:%s response: %s", compartmentId, response)
	backups := response.Items
	for _, i := range backups {
		fmt.Println(*i.Id, *i.DisplayName)
		return *i.DisplayName, nil
	}

	return "", nil
}

func DeleteBlockVolumeBackups(compartmentId string) (string, error) {
	coreClinet, clienrError := core.NewBlockstorageClientWithConfigurationProvider(common.DefaultConfigProvider())
	if clienrError != nil {
		return "", fmt.Errorf("Unable to create OCI clinet: %s", clienrError)
	}
	ctx := context.Background()

	glog.Info("List backup volumes for compartment: ", compartmentId)

	request := core.ListVolumeBackupsRequest{
		CompartmentId: &compartmentId,
	}
	response, err := coreClinet.ListVolumeBackups(ctx, request)
	if err != nil {
		return "", fmt.Errorf("Unable to list volume backup for compartment: %s, err:%s ", compartmentId, err)
	}

	fmt.Println("Volume backups for compartmentId:%s response: %s", compartmentId, response)
	backups := response.Items
	for _, i := range backups {
		fmt.Println(*i.Id, *i.DisplayName)

		request1 := core.DeleteVolumeBackupRequest{
			VolumeBackupId: i.Id,
		}

		_, err := coreClinet.DeleteVolumeBackup(ctx, request1)
		if err != nil {
			fmt.Println("Not Deleted: %s, err:%v ", *i.Id, err)
		}
		fmt.Println("Deleted: %s", *i.DisplayName)
	}
	return "", nil
}
