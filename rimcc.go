package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

var logger = shim.NewLogger("RIM")

type SimpleChaincode struct {
}


type CertificateDetails struct {
	AssetType        		string          `json:"docType"`
	CertificateID			string			`json:"certificateId"`
	CultivatorID			string			`json:"cultivatorId"`
	Name					string			`json:"name"`
	Length					string			`json:"length"`
	Breadth					string			`json:"breadth"`
	PlantInfo			    []string		`json:"plantInfo"`
	BudHarvestDtls			[]string		`json:"budHarvestDtls"`
	LeafHarvestDtls			[]string		`json:"leafHarvestDtls"`
	ValidTill				string			`json:"validTill"`
	IssuedOn				string			`json:"issuedOn"`

}

type PlantDetails struct {
	AssetType				string 			`json:"docType"`
	PlantID					string			`json:"plantId"`
	RemainingLeafCapacity	string			`json:"remainingLeafCapacity"`
	RemainingBudCapacity	string			`json:"remainingBudCapacity"`
	LeafHarvestCapacity		string			`json:"leafHarvestCapacity"`
	BudHarvestCapacity		string			`json:"budHarvestCapacity"`
	CertificateID			string			`json:"certificateId"`
	CultivatorID			string			`json:"cultivatorId"`


}


type HarvestDetails struct {
	AssetType        		string          `json:"docType"`
	HarvestID				string			`json:"harvestId"`
	HarvestType				string			`json:"harvestType"`
	LotCapacity				string			`json:"lotCapacity"`
	TestingStatus			string			`json:"testingStatus"`
	Grade					string			`json:"grade"`
	TestingRemarks			string			`json:"testingRemarks"`
	CreatedDate				string			`json:"createdDate"`
	DispatchedDate			string			`json:"dispatchedDate"`
	TestingDate				string			`json:"testingDate"`
	QALabID					string			`json:"qaLabId"`
	DestroyerQALabID		string			`json:"destroyerQaLabId"`
	CultivatorID			string			`json:"cultivatorId"`
	CertificateID			string			`json:"certificateId"`
	ExpiryDate				string			`json:"expiryDate"`
	DestroyedDate			string			`json:"destroyedDate"`
	DestroyedRemarks		string			`json:"destroyedRemarks"`
	HarvestPlantTags		[]string		`json:"harvestPlantTags"`
	HarvestPlantConsumed	[]string		`json:"harvestPlantConsumed"`
	LotInfo					[]string		`json:"lotInfo"`
	Status					string			`json:"status"`
	
}

type CultivatorDetails struct {
	AssetType        		string          `json:"docType"`
	CultivatorID			string			`json:"cultivatorId"`
	Name					string			`json:"name"`
	SSN						string			`json:"ssn"`
	Length					string			`json:"length"`
	Breadth					string			`json:"breadth"`
	LicenseClass			string			`json:"licenseClass"`
	Latitude				string			`json:"latitude"`
	Longitude				string			`json:"longitude"`
	Address					string			`json:"address"`
	State					string			`json:"state"`
	PhoneNumber				string			`json:"phoneNumber"`
	DelFlag					string			`json:"delFlag"`
	IsCertified				string			`json:"isCertified"`

}

type CultivatorResponse struct {
	CertificateAsset	string				`json:"certificateAsset"`
	PlantAsset			string				`json:"plantAsset"`
	HarvestAsset		string				`json:"harvestAsset"`
}

type LotDetails struct {
	AssetType			string				`json:"docType"`
	LotID				string				`json:"lotId"`
	LotType				string				`json:"lotType"`
	Grade				string				`json:"grade"`
	QALabID				string				`json:"qaLabId"`
	CreatedDate			string				`json:"createdDate"`
	DispatchedDate		string				`json:"dispatchedDate"`
	PackagerID			string				`json:"packagerId"`
	LotBatchID			string				`json:"lotBatchId"`
	TestingDate			string				`json:"testingDate"`
	TestingRemarks		string				`json:"testingRemarks"`
	Status				string				`json:"status"`
	PacketCapacity		string				`json:"packetCapacity"`
	PacketInfo			[]string			`json:"packetInfo"`
//PacketCapacity will be set during createLot,needed for creating packet assests according to their
//types. PacketInfo needed to store the array of packetIds created from that particular lot.
}

type PacketDetails struct {
	AssetType			string				`json:"docType"`
	PacketCreationDate	string				`json:"packetCreationDate"`
	PacketID			string				`json:"packetId"`
	LotID				string				`json:"lotId"`
	Type				string				`json:"type"`
	Grade				string				`json:"grade"`
	PacketBatchID		string				`json:"packetBatchId"`
	PackagerID			string				`json:"packagerId"`
	PatientID			string				`json:"patientId"`
	Prescription		string				`json:"prescription"`
	Dispatched			string				`json:"dispatched"`
	IssueDate			string				`json:"issueDate"`
}

type PurchaseOrderDetails struct {
	AssetType			string				`json:"docType"`
	PurchaseOrderID		string				`json:"purchaseOrderId"`
	Type				string				`json:"type"`
	Grade				string				`json:"grade"`
	Quantity			string				`json:"quantity"`
	OrderDescription	string				`json:"orderDescription"`
	PackagerID			string				`json:"packagerId"`
	PacketInfo			[]string			`json:"packetInfo"`
	DispensaryID		string				`json:"dispensaryId"`
	OrderDate			string				`json:"orderDate"`
	DispatchedDate		string				`json:"dispatchedDate"`
}

type PacketResponse struct {
	PurchaseOrderAsset	string				`json:"purchaseOrderAsset"`
	PacketAsset			string				`json:"packetAsset"`
}

type PacketLotResponse struct {
	LotAsset			string				`json:"lotAsset"`
	PacketAsset			string				`json:"packetAsset"`
}

type PatientDetails struct {
	AssetType			string				`json:"docType"`
	PatientID			string				`json:"patientId"`
	Type				string				`json:"type"`
	Grade				string				`json:"grade"`
	PhoneNumber			string				`json:"phoneNumber"`
	Prescription		string				`json:"prescription"`
	PacketInfo			[]string			`json:"packetInfo"`
	DispensaryID		string				`json:"dispensaryId"`

}

type PatientdtByDisPensary struct{
	AssetType			string				`json:"docType"`
	DispensaryID		string				`json:"dispensaryId"`
	PatientDetails      []PatientDetails    `json:"patientDetails"`
}


type PatientPrescription struct{
	AssetType				string				`json:"docType"`
	Date					string				`json:"date"`
	Doses					string				`json:"doses"`
	Type					string				`json:"type"`
	Grade					string				`json:"grade"`
	Status					string				`json:"status"`
	DoctorID				string				`json:"doctorid"`
	Prescription			string				`json:"prescription"`
}

type Patient struct {
	AssetType				string					`json:"docType"`
	PatientId				string					`json:"patientId"`
	Age						string					`json:"Age"`
	SSN				 		string					`json:"SSN"`
	PhoneNo					string					`json:"phoneNo"`
	Name		      		string					`json:"name"`
	Dispensary				string					`json:"dispensary"`
	Patient_Prescription	[]PatientPrescription	`json:"patient_Prescription"`

}

type PurchaseResponse struct{
	PurchaseOrderAsset	string			`json:"purchaseOrderAsset"`
	PacketAsset			string			`json:"packetAsset"`
}

type PatientResponse struct{
	PacketAsset			string			`json:"packetAsset"` 
	PatientAsset		string			`json:"patientAsset"`
}

type TrackResponse struct {
	PacketAsset 		string				`json:"packetAsset"`
	PatientAsset		string				`json:"patientAsset"`
	PurchaseOrderAsset  string				`json:"purchaseOrderAsset"`
	LotAsset 			string				`json:"lotAsset"`
	HarvestAsset		string				`json:"harvestAsset"`
	PlantAsset			string				`json:"plantAsset"`
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {

	logger.Info("########### rimcc Init ###########")
	return shim.Success(nil)
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("########### RIM Invoke ###########")

	function, args := stub.GetFunctionAndParameters()

	if function == "init" {
		return t.Init(stub)
	}

	if function == "cultivatorRegistration" {
		return t.cultivatorRegistration(stub, args)
	}

	if function == "certificateIssue" {
		return t.certificateIssue(stub, args)
	}

	if function == "viewAllCultivatorDtls" {
		return t.viewAllCultivatorDtls(stub)
	}
	
	if function == "viewCultivatorDtlsById" {
		return t.viewCultivatorDtlsById(stub, args)
	}
	
	if function == "viewCertificateDtlsByCultivatorId" {
		return t.viewCertificateDtlsByCultivatorId(stub, args)
	}
	
	if function == "updateCultivatorDtls" {
		return t.updateCultivatorDtls(stub, args)
	}

	if function == "deleteCultivatorDtls" {
		return t.deleteCultivatorDtls(stub, args)
	}
	
	if function == "dispatchHarvest" {
		return t.dispatchHarvest(stub, args)
	}

	if function == "viewHarvestDtlsById" {
		return t.viewHarvestDtlsById(stub, args)
	}
	
	if function == "viewHarvestDtlsByCultivatorId" {
		return t.viewHarvestDtlsByCultivatorId(stub, args)
	}

	if function == "fetchDataForCultivator" {
		return t.fetchDataForCultivator(stub, args)
	}
	
	if function == "updateHarvestDtls" {
		return t.updateHarvestDtls(stub, args)
	}

	if function == "viewHarvestDtlsByQALabId" {
		return t.viewHarvestDtlsByQALabId(stub, args)
	}

	if function == "viewHarvestDtlsByGrade" {
		return t.viewHarvestDtlsByGrade(stub, args)
	}

	if function == "destroyHarvest" {
		return t.destroyHarvest(stub, args)
	}

	if function == "splitHarvest" {
		return t.splitHarvest(stub, args)
	}

	if function == "createLot" {
		return t.createLot(stub, args)
	}

	if function == "viewLotDtls" {
		return t.viewLotDtls(stub)
	}
	
	if function == "dispatchLot" {
		return t.dispatchLot(stub, args)
	}
	
	if function == "viewLotDtlsByPackagerId" {
		return t.viewLotDtlsByPackagerId(stub, args)
	}

	if function == "createPacket" {
		return t.createPacket(stub, args)
	}

	if function == "viewPacketDtlsByLotId" {
		return t.viewPacketDtlsByLotId(stub, args)
	}

	if function == "viewPurchaseOrderByPackagerId" {
		return t.viewPurchaseOrderByPackagerId(stub, args)
	}

	if function == "dispatchPacket" {
		return t.dispatchPacket(stub, args)
	}

	if function == "viewDispatchedPurchaseOrder" {
		return t.viewDispatchedPurchaseOrder(stub, args)
	}

	if function == "createPurchaseOrder" {
		return t.createPurchaseOrder(stub, args)
	}
	
	if function == "viewPurchaseOrderByDispensaryId" {
		return t.viewPurchaseOrderByDispensaryId(stub, args)
	}

	if function == "viewPacketByDispensaryId" {
		return t.viewPacketByDispensaryId(stub, args)
	}

	if function == "issuePacket" {
		return t.issuePacket(stub, args)
	}

	if function == "viewPacketDtlsByPacketId" {
		return t.viewPacketDtlsByPacketId(stub, args)
	}

	if function == "viewPacketDtlsByPatientId" {
		return t.viewPacketDtlsByPatientId(stub, args)
	}
	if function == "viewPacketDtlsByIssueDate" {
		return t.viewPacketDtlsByIssueDate(stub, args)
	}

	if function == "trackPackageByPacketId" {
		return t.trackPackageByPacketId(stub, args)
	}

	if function== "addPatient" {
		return t.addPatient(stub,args)
	}

	if function== "getPatientDetails" {
		return t.getPatientDetails(stub)
	}
	


	logger.Errorf("Unknown action, check the first argument, must be one of 'cultivatorRegistration', 'certificateIssue', 'viewAllCultivatorDtls', 'viewCultivatorDtlsById', 'viewCertificateDtlsByCultivatorId', 'updateCultivatorDtls', 'deleteCultivatorDtls', 'dispatchHarvest', 'viewHarvestDtlsById', 'viewHarvestDtlsByCultivatorId', 'fetchDataForCultivator', 'updateHarvestDtls', 'viewHarvestDtlsByQALabId', 'viewHarvestDtlsByGrade', 'destroyHarvest', 'splitHarvest', 'createLot', 'viewLotDtls', 'dispatchLot', 'viewLotDtlsByPackagerId','createPacket','viewPacketDtlsByLotId','viewPurchaseOrderByPackagerId','dispatchPacket','viewDispatchedPurchaseOrder','createPurchaseOrder','viewPurchaseOrderByDispensaryId','viewPacketByDispensaryId','issuePacket','viewPacketDtlsByPacketId','viewPacketDtlsByPatientId','trackPackageByPacketId','viewPacketDtlsByIssueDate','addPatient','getPatientDetails'. But got: %v", function)
	return shim.Error(fmt.Sprintf("Unknown action, check the first argument, must be one of 'cultivatorRegistration', 'certificateIssue', 'viewAllCultivatorDtls', 'viewCultivatorDtlsById', 'viewCertificateDtlsByCultivatorId', 'updateCultivatorDtls', 'deleteCultivatorDtls','dispatchHarvest', 'viewHarvestDtlsById', 'viewHarvestDtlsByCultivatorId' , 'fetchDataForCultivator', 'updateHarvestDtls', 'viewHarvestDtlsByQALabId', 'viewHarvestDtlsByGrade', 'destroyHarvest', 'splitHarvest', 'createLot', 'viewLotDtls', 'dispatchLot', 'viewLotDtlsByPackagerId','createPacket','viewPacketDtlsByLotId','viewPurchaseOrderByPackagerId','dispatchPacket','viewDispatchedPurchaseOrder','createPurchaseOrder','viewPurchaseOrderByDispensaryId','viewPacketByDispensaryId','issuePacket''viewPacketDtlsByPacketId','viewPacketDtlsByPatientId','trackPackageByPacketId','viewPacketDtlsByIssueDate','addPatient','getPatientDetails'. But got: %v", function))

}

func (t *SimpleChaincode) cultivatorRegistration(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 11 {
		return shim.Error("Incorrect number of arguments. Expecting 11")
	}

	cultivatorIdIterator, err := stub.GetQueryResult("{\"selector\":{\"docType\":\"CultivatorDetails\",\"cultivatorID\":\"" + args[0] + "\"}}")
	if err != nil {
		return shim.Error(err.Error())
	}

	if cultivatorIdIterator.HasNext() {
		return shim.Success([]byte("Cultivator Id " + args[0] + "already exist."))
	} else {

			ssnIdIterator, err := stub.GetQueryResult("{\"selector\":{\"docType\":\"CultivatorDetails\",\"ssn\":\"" + args[2] + "\"}}")
			if err != nil {
				return shim.Error(err.Error())
			}

			if ssnIdIterator.HasNext() {
				return shim.Success([]byte("SSN exists"))
			} else {
						var cultivatorDetail = CultivatorDetails{
							AssetType:   "CultivatorDetails",
							CultivatorID:     args[0],
							Name:   		  args[1],
							SSN:    	  	  args[2],
							Length:	  	  	  args[3],
							Breadth:		  args[4],
							LicenseClass:     args[5],
							Latitude:		  args[6],
							Longitude:		  args[7],
							Address:     	  args[8],
							State:			  args[9],
							PhoneNumber:	  args[10],
							DelFlag:          "N",
							IsCertified:  	  "N"}
						cultivatorDetailAsBytes, _ := json.Marshal(cultivatorDetail)
						err = stub.PutState(args[0], cultivatorDetailAsBytes)
						if err != nil {
							return shim.Error("Error while put state of cultivatorDetail for cultivator registration.")
						}
						return shim.Success([]byte("CultivatorId " + args[0] + " created Successfully."))
					}
			}
	
	return shim.Success(nil)
}

func (t *SimpleChaincode) updateCultivatorDtls(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 10 {
		return shim.Error("Incorrect number of arguments. Expecting 10")
	}

	cultivatorIdIterator, err := stub.GetQueryResult("{\"selector\":{\"docType\":\"CultivatorDetails\",\"cultivatorId\":\"" + args[0] + "\"}}")
	if err != nil {
		return shim.Error(err.Error())
	}

	if cultivatorIdIterator.HasNext() {
		var err error
		resultQuery, resultError := cultivatorIdIterator.Next()
		if resultError != nil {
			return shim.Error(resultError.Error())
		}
		var cultivatorDetails CultivatorDetails
		unMarshalError := json.Unmarshal(resultQuery.Value, &cultivatorDetails)
		if unMarshalError != nil {
			return shim.Error(unMarshalError.Error())
		}

		var cultivatorDetail = CultivatorDetails{
			AssetType:   	  	"CultivatorDetails",
			CultivatorID:     	cultivatorDetails.CultivatorID,
			Name:   		  	args[1],
			SSN:  	  	  	  	cultivatorDetails.SSN,
			Length:	  	  	  	args[2],
			Breadth:		  	args[3],
			LicenseClass:     	args[4],
			Latitude:		  	args[5],
			Longitude:		  	args[6],
			Address:     	  	args[7],
			State:			  	args[8],
			PhoneNumber:	  	args[9],
			DelFlag:          	cultivatorDetails.DelFlag,
			IsCertified:  	  	cultivatorDetails.IsCertified}
			
			cultivatorDetailAsBytes, _ := json.Marshal(cultivatorDetail)
		err = stub.PutState(args[0], cultivatorDetailAsBytes)
		if err != nil {
			return shim.Error("Error while put state to update cultivatorDetail")
		}
	} else {
		return shim.Error("Cutivator Id" + args[0] + " doesn't exist. Register the Cultivator first.")
	}
	return shim.Success([]byte("CultivatorId " + args[0] + " updated successfully."))
}

func (t *SimpleChaincode) deleteCultivatorDtls(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	logger.Info("Arguments:", args[0])
	cultivatorIdIterator, err := stub.GetQueryResult("{\"selector\":{\"docType\":\"CultivatorDetails\",\"cultivatorId\":\"" + args[0] + "\"}}")
	if err != nil {
		return shim.Error(err.Error())
	}

	if cultivatorIdIterator.HasNext() {
		var err error
		resultQuery, resultError := cultivatorIdIterator.Next()
		if resultError != nil {
			return shim.Error(resultError.Error())
		}
		var cultivatorDetails CultivatorDetails
		unMarshalError := json.Unmarshal(resultQuery.Value, &cultivatorDetails)
		if unMarshalError != nil {
			return shim.Error(unMarshalError.Error())
		}
		logger.Info("cultivator details:", cultivatorDetails)
		var cultivatorDetail = CultivatorDetails{
			AssetType:   		"CultivatorDetails",
			CultivatorID:     	cultivatorDetails.CultivatorID,
			Name:   		  	cultivatorDetails.Name,
			SSN:  	  	  	  	cultivatorDetails.SSN,
			Length:	  	  	  	cultivatorDetails.Length,
			Breadth:		  	cultivatorDetails.Breadth,
			LicenseClass:     	cultivatorDetails.LicenseClass,
			Latitude:		  	cultivatorDetails.Latitude,
			Longitude:		  	cultivatorDetails.Longitude,
			Address:     	  	cultivatorDetails.Address,
			State:			  	cultivatorDetails.State,
			PhoneNumber:	  	cultivatorDetails.PhoneNumber,
			DelFlag:          	"Y",
			IsCertified:  	  	cultivatorDetails.IsCertified}
			logger.Info("cultivator details updated:", cultivatorDetail)
		cultivatorDetailAsBytes, _ := json.Marshal(cultivatorDetail)
		err = stub.PutState(args[0], cultivatorDetailAsBytes)
		if err != nil {
			return shim.Error("Error while updating put state of cultivator details for certificate issue.")
		}
	} else {
		return shim.Error("Cultivator Id" + args[0] + " doesn't exist. Register the Cultivator first.")
	}
	return shim.Success([]byte("CultivatorId " + args[0] + " deleted successfully."))
}

func (t *SimpleChaincode) certificateIssue(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 10 {
		return shim.Error("Incorrect number of arguments. Expecting 10")
	}
	logger.Info("Arguments:", args[5])
	certificateIdIterator, err := stub.GetQueryResult("{\"selector\":{\"docType\":\"CertificateDetails\",\"certificateId\":\"" + args[0] + "\"}}")
	if err != nil {
		return shim.Error(err.Error())
	}

	if certificateIdIterator.HasNext() {
		return shim.Success([]byte("Certificate Id" + args[0] + "already exist."))
	}else {

		var plantArray[] string
		plantArrayAsBytes:=[]byte(args[5])
		unMarshalError :=json.Unmarshal(plantArrayAsBytes,&plantArray)
		if unMarshalError != nil {
			return shim.Error(unMarshalError.Error())
		}
		for i := 0; i < len(plantArray); i++ {
			var plantDetails = PlantDetails{
			AssetType:   			"PlantDetails",
			PlantID: 	 			plantArray[i],
			RemainingLeafCapacity: 	"2",
			RemainingBudCapacity:	"1",
			LeafHarvestCapacity:	"2",
			BudHarvestCapacity:		"1",
			CertificateID:    		args[0],
			CultivatorID:     		args[1]}
			plantDetailsAsBytes, _ := json.Marshal(plantDetails)
			err = stub.PutState(plantArray[i], plantDetailsAsBytes)
			if err != nil {
				return shim.Error(err.Error())
			}
		}


		var budArray[] string
		budArrayAsBytes:=[]byte(args[8])
		unMarshalBudError :=json.Unmarshal(budArrayAsBytes,&budArray)
		if unMarshalBudError != nil {
			return shim.Error(unMarshalError.Error())
		}
		
		for i := 0; i < len(budArray); i++ {
			var harvestDetailsBud = HarvestDetails{
			AssetType:   	"HarvestDetails",
			HarvestID: 	 	budArray[i],
			HarvestType: 	"Bud",
			CreatedDate:	args[7],
			LotCapacity:	"10",
			CultivatorID:	args[1],
			CertificateID:	args[0],
			Status:			"C"}
			
			harvestDetailsBudAsBytes, _ := json.Marshal(harvestDetailsBud)
			err = stub.PutState(budArray[i], harvestDetailsBudAsBytes)
			if err != nil {
				return shim.Error("Error while put state of harvestDetails of bud for certificate issue.")
			}
		}
		
		var leafArray[] string
		leafArrayAsBytes:=[]byte(args[9])
		unMarshalLeafError :=json.Unmarshal(leafArrayAsBytes,&leafArray)
		if unMarshalLeafError != nil {
			return shim.Error(unMarshalError.Error())
		}
		
		for i := 0; i < len(leafArray); i++ {
			var harvestDetailsLeaf = HarvestDetails{
			AssetType:   	"HarvestDetails",
			HarvestID: 	  	leafArray[i],
			HarvestType: 	"Leaf",
			CreatedDate:	args[7],
			LotCapacity:	"20",
			CultivatorID:	args[1],
			CertificateID:	args[0],
			Status:			"C"}
			
			harvestDetailsLeafAsBytes, _ := json.Marshal(harvestDetailsLeaf)
			err = stub.PutState(leafArray[i], harvestDetailsLeafAsBytes)
			if err != nil {
				return shim.Error("Error while put state of harvestDetails of leaf for certificate issue.")
			}
		}
		
		var certificateDetails = CertificateDetails{
			AssetType:   		"CertificateDetails",
			CertificateID:    	args[0],
			CultivatorID:     	args[1],
			Name:				args[2],
			Length:				args[3],
			Breadth:			args[4],
			PlantInfo: 			plantArray,
			ValidTill:	  	  	args[6],
			IssuedOn:     		args[7],
			BudHarvestDtls:		budArray,
			LeafHarvestDtls:	leafArray}
		certificateDetailsAsBytes, _ := json.Marshal(certificateDetails)
		err = stub.PutState(args[0], certificateDetailsAsBytes)
		if err != nil {
			return shim.Error("Error while put state of certificateDetails for certificate issue.")
		}
		cultivatorIdIterator, err := stub.GetQueryResult("{\"selector\":{\"docType\":\"CultivatorDetails\",\"cultivatorId\":\"" + args[1] + "\"}}")
		if err != nil {
			return shim.Error(err.Error())
		}

		if cultivatorIdIterator.HasNext() {
			var err error
			resultQuery, resultError := cultivatorIdIterator.Next()
			if resultError != nil {
				return shim.Error(resultError.Error())
			}
			var cultivatorDetails CultivatorDetails
			unMarshalError := json.Unmarshal(resultQuery.Value, &cultivatorDetails)
			if unMarshalError != nil {
				return shim.Error(unMarshalError.Error())
			}

			var cultivatorDetail = CultivatorDetails{
				AssetType:   		"CultivatorDetails",
				CultivatorID:     	cultivatorDetails.CultivatorID,
				Name:   		  	cultivatorDetails.Name,
				SSN:  	  	  	  	cultivatorDetails.SSN,
				Length:	  	  	  	cultivatorDetails.Length,
				Breadth:		  	cultivatorDetails.Breadth,
				LicenseClass:     	cultivatorDetails.LicenseClass,
				Latitude:		  	cultivatorDetails.Latitude,
				Longitude:		  	cultivatorDetails.Longitude,
				Address:     	  	cultivatorDetails.Address,
				State:			  	cultivatorDetails.State,
				PhoneNumber:	  	cultivatorDetails.PhoneNumber,
				DelFlag:          	cultivatorDetails.DelFlag,
				IsCertified:  	  	"Y"}
			cultivatorDetailAsBytes, _ := json.Marshal(cultivatorDetail)
			err = stub.PutState(args[1], cultivatorDetailAsBytes)
			if err != nil {
				return shim.Error("Error while updating put state of cultivator details for certificate issue.")
			}
			return shim.Success([]byte("CertificateId " + args[0] + " created successfully."))
		}
	}
	return shim.Success(nil)
}

func (t *SimpleChaincode) viewAllCultivatorDtls(stub shim.ChaincodeStubInterface) pb.Response {


	cultivatorIterator, err := stub.GetQueryResult("{\"selector\":{\"docType\":\"CultivatorDetails\",\"delFlag\":\"N\"}}")
	if err != nil {
		return shim.Error(err.Error())
	}

	defer cultivatorIterator.Close()
	// buffer is a JSON array containing QueryResults
	var cultivatorBuffer bytes.Buffer
	cultivatorBuffer.WriteString("[")

	cultivatorAlreadyWritten := false
	for cultivatorIterator.HasNext() {
		queryResponse, err := cultivatorIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if cultivatorAlreadyWritten == true {
			cultivatorBuffer.WriteString(",")
		}
		cultivatorBuffer.WriteString("{\"Key\":")
		cultivatorBuffer.WriteString("\"")
		cultivatorBuffer.WriteString(queryResponse.Key)
		cultivatorBuffer.WriteString("\"")

		cultivatorBuffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		cultivatorBuffer.WriteString(string(queryResponse.Value))
		cultivatorBuffer.WriteString("}")
		cultivatorAlreadyWritten = true
	}
	cultivatorBuffer.WriteString("]")

	return shim.Success(cultivatorBuffer.Bytes())
}


func (t *SimpleChaincode) viewCultivatorDtlsById(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	cultivatorIdIterator, err := stub.GetQueryResult("{\"selector\":{\"docType\":\"CultivatorDetails\",\"cultivatorId\":\"" + args[0] + "\",\"delFlag\":\"N\"}}")
	if err != nil {
		return shim.Error(err.Error())
	}

	defer cultivatorIdIterator.Close()
	// buffer is a JSON array containing QueryResults
	var cultivatorIdBuffer bytes.Buffer
	cultivatorIdBuffer.WriteString("[")

	cultivatorIdAlreadyWritten := false
	for cultivatorIdIterator.HasNext() {
		queryResponse, err := cultivatorIdIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if cultivatorIdAlreadyWritten == true {
			cultivatorIdBuffer.WriteString(",")
		}
		cultivatorIdBuffer.WriteString("{\"Key\":")
		cultivatorIdBuffer.WriteString("\"")
		cultivatorIdBuffer.WriteString(queryResponse.Key)
		cultivatorIdBuffer.WriteString("\"")

		cultivatorIdBuffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		cultivatorIdBuffer.WriteString(string(queryResponse.Value))
		cultivatorIdBuffer.WriteString("}")
		cultivatorIdAlreadyWritten = true
	}
	cultivatorIdBuffer.WriteString("]")

	return shim.Success(cultivatorIdBuffer.Bytes())
}

func (t *SimpleChaincode) viewCertificateDtlsByCultivatorId(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	certificateIdIterator, err := stub.GetQueryResult("{\"selector\":{\"docType\":\"CertificateDetails\",\"cultivatorId\":\""+ args[0] + "\"}}")
	if err != nil {
		return shim.Error(err.Error())
	}

	defer certificateIdIterator.Close()
	// buffer is a JSON array containing QueryResults
	var certificateIdBuffer bytes.Buffer
	certificateIdBuffer.WriteString("[")

	certificateIdAlreadyWritten := false
	for certificateIdIterator.HasNext() {
		queryResponse, err := certificateIdIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if certificateIdAlreadyWritten == true {
			certificateIdBuffer.WriteString(",")
		}
		certificateIdBuffer.WriteString("{\"Key\":")
		certificateIdBuffer.WriteString("\"")
		certificateIdBuffer.WriteString(queryResponse.Key)
		certificateIdBuffer.WriteString("\"")

		certificateIdBuffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		certificateIdBuffer.WriteString(string(queryResponse.Value))
		certificateIdBuffer.WriteString("}")
		certificateIdAlreadyWritten = true
	}
	certificateIdBuffer.WriteString("]")

	return shim.Success(certificateIdBuffer.Bytes())
	
}


func (t *SimpleChaincode) dispatchHarvest(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 8 {
		return shim.Error("Incorrect number of arguments. Expecting 8")
	}
		var plantHarvestArray[] string
		var harvestPlantTagsArray [] string
		var harvestPlantConsumedArray [] string
		plantHarvestArrayAsbytes:=[]byte(args[4])
		unMarshalplantHarvestArrayError:=json.Unmarshal(plantHarvestArrayAsbytes, &plantHarvestArray)
		if unMarshalplantHarvestArrayError!=nil {
			return shim.Error(unMarshalplantHarvestArrayError.Error())
		}
	for i := 0; i < len(plantHarvestArray); i++ {
		
		plantIdIterator, err := stub.GetQueryResult("{\"selector\":{\"docType\":\"PlantDetails\",\"plantId\":\""+ plantHarvestArray[i] + "\"}}")
	if err != nil {
		return shim.Error(err.Error())
	}

	if plantIdIterator.HasNext() {
		var err error
		resultQueryPlant, resultError := plantIdIterator.Next()
		if resultError != nil {
			return shim.Error(resultError.Error())
		}

		
		var plantDetails PlantDetails
		unMarshalPlantError:=json.Unmarshal(resultQueryPlant.Value, &plantDetails)
		if unMarshalPlantError!=nil {
			return shim.Error(unMarshalPlantError.Error())
		}

		if args[1]=="Bud" {
			var plantDetailBud = PlantDetails{
				AssetType:				"PlantDetails",
				PlantID:				plantDetails.PlantID,
				RemainingLeafCapacity:	plantDetails.RemainingLeafCapacity,
				RemainingBudCapacity:	plantHarvestArray[i+1],
				LeafHarvestCapacity:	plantDetails.LeafHarvestCapacity,
				BudHarvestCapacity:		plantDetails.BudHarvestCapacity,
				CertificateID:			plantDetails.CertificateID,
				CultivatorID:			plantDetails.CultivatorID}
				plantDetailBudAsBytes, _ := json.Marshal(plantDetailBud)
				err = stub.PutState(plantHarvestArray[i], plantDetailBudAsBytes)
				if err != nil {
					return shim.Error("Error while put state to update harvestDetail")
				}	
		} else {
			var plantDetailLeaf = PlantDetails{
				AssetType:				"PlantDetails",
				PlantID:				plantDetails.PlantID,
				RemainingLeafCapacity:	plantHarvestArray[i+1],
				RemainingBudCapacity:	plantDetails.RemainingBudCapacity,
				LeafHarvestCapacity:	plantDetails.LeafHarvestCapacity,
				BudHarvestCapacity:		plantDetails.BudHarvestCapacity,
				CertificateID:			plantDetails.CertificateID,
				CultivatorID:			plantDetails.CultivatorID}
				plantDetailBudAsBytes, _ := json.Marshal(plantDetailLeaf)
				err = stub.PutState(plantHarvestArray[i], plantDetailBudAsBytes)
				if err != nil {
					return shim.Error("Error while put state to update harvestDetail")
				}	
			}
			harvestPlantTagsArray= append(harvestPlantTagsArray, plantHarvestArray[i])
			harvestPlantConsumedArray= append(harvestPlantConsumedArray, plantHarvestArray[i+2])
		}
		i+=2
	}
	harvestIdIterator, err := stub.GetQueryResult("{\"selector\":{\"docType\":\"HarvestDetails\",\"harvestId\":\"" + args[0] + "\",\"harvestType\":\"" + args[1] + "\"}}")
	if err != nil {
		return shim.Error(err.Error())
	}

	if harvestIdIterator.HasNext() {
		var err error
		resultQuery, resultError := harvestIdIterator.Next()
		if resultError != nil {
			return shim.Error(resultError.Error())
		}
		var harvestDetails HarvestDetails
		unMarshalError := json.Unmarshal(resultQuery.Value, &harvestDetails)
		if unMarshalError != nil {
			return shim.Error(unMarshalError.Error())
		}
		

		var harvestDetail = HarvestDetails{
			AssetType:   	  		"HarvestDetails",
			HarvestID:     	  		harvestDetails.HarvestID,
			HarvestType:   	  		harvestDetails.HarvestType,
			LotCapacity:			harvestDetails.LotCapacity,
			TestingStatus:	  		args[2],
			CreatedDate:			harvestDetails.CreatedDate,
			DispatchedDate:			args[6],
			TestingDate:			args[5],
			QALabID:				args[3],
			CultivatorID:			harvestDetails.CultivatorID,
			CertificateID:			harvestDetails.CertificateID,
			ExpiryDate:      		args[7],
			HarvestPlantTags:   	harvestPlantTagsArray,
			HarvestPlantConsumed: 	harvestPlantConsumedArray,
			Status:					"D"}
			harvestDetailAsBytes, _ := json.Marshal(harvestDetail)
		err = stub.PutState(args[0], harvestDetailAsBytes)
		if err != nil {
			return shim.Error("Error while put state to update harvestDetail")
		}
	} else {
		return shim.Error("Harvest Id" + args[0] + " doesn't exist. Register the Harvest first.")
	}
	return shim.Success([]byte("HarvestId " + args[0] + " dispatched successfully."))
}


func (t *SimpleChaincode) viewHarvestDtlsById(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	harvestIdIterator, err := stub.GetQueryResult("{\"selector\":{\"docType\":\"HarvestDetails\",\"harvestId\":\""+ args[0] + "\"}}")
	if err != nil {
		return shim.Error(err.Error())
	}

	defer harvestIdIterator.Close()
	// buffer is a JSON array containing QueryResults
	var harvestIdBuffer bytes.Buffer
	harvestIdBuffer.WriteString("[")

	harvestIdAlreadyWritten := false
	for harvestIdIterator.HasNext() {
		queryResponse, err := harvestIdIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if harvestIdAlreadyWritten == true {
			harvestIdBuffer.WriteString(",")
		}
		harvestIdBuffer.WriteString("{\"Key\":")
		harvestIdBuffer.WriteString("\"")
		harvestIdBuffer.WriteString(queryResponse.Key)
		harvestIdBuffer.WriteString("\"")

		harvestIdBuffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		harvestIdBuffer.WriteString(string(queryResponse.Value))
		harvestIdBuffer.WriteString("}")
		harvestIdAlreadyWritten = true
	}
	harvestIdBuffer.WriteString("]")

	return shim.Success(harvestIdBuffer.Bytes())
	
}


func (t *SimpleChaincode) viewHarvestDtlsByCultivatorId(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	harvestIdIterator, err := stub.GetQueryResult("{\"selector\":{\"docType\":\"HarvestDetails\",\"cultivatorId\":\"" + args[0] + "\",\"status\":\"D\"}}")
	if err != nil {
		return shim.Error(err.Error())
	}

	defer harvestIdIterator.Close()
	// buffer is a JSON array containing QueryResults
	var harvestIdBuffer bytes.Buffer
	harvestIdBuffer.WriteString("[")

	harvestIdAlreadyWritten := false
	for harvestIdIterator.HasNext() {
		queryResponse, err := harvestIdIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if harvestIdAlreadyWritten == true {
			harvestIdBuffer.WriteString(",")
		}
		harvestIdBuffer.WriteString("{\"Key\":")
		harvestIdBuffer.WriteString("\"")
		harvestIdBuffer.WriteString(queryResponse.Key)
		harvestIdBuffer.WriteString("\"")

		harvestIdBuffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		harvestIdBuffer.WriteString(string(queryResponse.Value))
		harvestIdBuffer.WriteString("}")
		harvestIdAlreadyWritten = true
	}
	harvestIdBuffer.WriteString("]")

	return shim.Success(harvestIdBuffer.Bytes())	
}


func (t *SimpleChaincode) fetchDataForCultivator(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	certificateIdIterator, err := stub.GetQueryResult("{\"selector\":{\"docType\":\"CertificateDetails\",\"cultivatorId\":\""+ args[0] + "\"}}")
	if err != nil {
		return shim.Error(err.Error())
	}

	defer certificateIdIterator.Close()
	// buffer is a JSON array containing QueryResults
	var certificateIdBuffer bytes.Buffer
	certificateIdBuffer.WriteString("[")

	certificateIdAlreadyWritten := false
	for certificateIdIterator.HasNext() {
		queryResponse, err := certificateIdIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if certificateIdAlreadyWritten == true {
			certificateIdBuffer.WriteString(",")
		}
		certificateIdBuffer.WriteString("{\"Key\":")
		certificateIdBuffer.WriteString("\"")
		certificateIdBuffer.WriteString(queryResponse.Key)
		certificateIdBuffer.WriteString("\"")

		certificateIdBuffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		certificateIdBuffer.WriteString(string(queryResponse.Value))
		certificateIdBuffer.WriteString("}")
		certificateIdAlreadyWritten = true
	}
	certificateIdBuffer.WriteString("]")
	
	plantIdIterator, err := stub.GetQueryResult("{\"selector\":{\"docType\":\"PlantDetails\",\"cultivatorId\":\""+ args[0] + "\"}}")
	if err != nil {
		return shim.Error(err.Error())
	}

	defer plantIdIterator.Close()
	// buffer is a JSON array containing QueryResults
	var plantIdBuffer bytes.Buffer
	plantIdBuffer.WriteString("[")

	plantIdAlreadyWritten := false
	for plantIdIterator.HasNext() {
		queryResponse, err := plantIdIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if plantIdAlreadyWritten == true {
			plantIdBuffer.WriteString(",")
		}
		plantIdBuffer.WriteString("{\"Key\":")
		plantIdBuffer.WriteString("\"")
		plantIdBuffer.WriteString(queryResponse.Key)
		plantIdBuffer.WriteString("\"")

		plantIdBuffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		plantIdBuffer.WriteString(string(queryResponse.Value))
		plantIdBuffer.WriteString("}")
		plantIdAlreadyWritten = true
	}
	plantIdBuffer.WriteString("]")

	harvestIdIterator, err := stub.GetQueryResult("{\"selector\":{\"docType\":\"HarvestDetails\",\"cultivatorId\":\"" + args[0] + "\",\"status\":\"C\"}}")
	if err != nil {
		return shim.Error(err.Error())
	}

	defer harvestIdIterator.Close()
	// buffer is a JSON array containing QueryResults
	var harvestIdBuffer bytes.Buffer
	harvestIdBuffer.WriteString("[")

	harvestIdAlreadyWritten := false
	for harvestIdIterator.HasNext() {
		queryResponse, err := harvestIdIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if harvestIdAlreadyWritten == true {
			harvestIdBuffer.WriteString(",")
		}
		harvestIdBuffer.WriteString("{\"Key\":")
		harvestIdBuffer.WriteString("\"")
		harvestIdBuffer.WriteString(queryResponse.Key)
		harvestIdBuffer.WriteString("\"")

		harvestIdBuffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		harvestIdBuffer.WriteString(string(queryResponse.Value))
		harvestIdBuffer.WriteString("}")
		harvestIdAlreadyWritten = true
	}
	harvestIdBuffer.WriteString("]")

	var res = CultivatorResponse{CertificateAsset: certificateIdBuffer.String(), PlantAsset: plantIdBuffer.String(), HarvestAsset: harvestIdBuffer.String()}
	finalResponse, marshalerr := json.Marshal(res)
	if marshalerr != nil {
		shim.Error(marshalerr.Error())
	}
	return shim.Success(finalResponse)
}


func (t *SimpleChaincode) updateHarvestDtls(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 6 {
		return shim.Error("Incorrect number of arguments. Expecting 6")
	}

	harvestIdIterator, err := stub.GetQueryResult("{\"selector\":{\"docType\":\"HarvestDetails\",\"harvestId\":\"" + args[0] + "\"}}")
	if err != nil {
		return shim.Error(err.Error())
	}

	if harvestIdIterator.HasNext() {
		var err error
		resultQuery, resultError := harvestIdIterator.Next()
		if resultError != nil {
			return shim.Error(resultError.Error())
		}
		var harvestDetails HarvestDetails
		unMarshalError := json.Unmarshal(resultQuery.Value, &harvestDetails)
		if unMarshalError != nil {
			return shim.Error(unMarshalError.Error())
		}
		
		var harvestDetail = HarvestDetails{
			AssetType:				"HarvestDetails",     		
			HarvestID:				harvestDetails.HarvestID,
			HarvestType:			harvestDetails.HarvestType,
			LotCapacity:			harvestDetails.LotCapacity,
			TestingStatus:			args[1],
			Grade:					args[2],
			TestingRemarks:			args[3],
			CreatedDate:			harvestDetails.CreatedDate,
			DispatchedDate:			harvestDetails.DispatchedDate,
			TestingDate:			args[5],
			QALabID:				args[4],
			CultivatorID:			harvestDetails.CultivatorID,
			CertificateID:			harvestDetails.CertificateID,
			ExpiryDate:				harvestDetails.ExpiryDate,
			HarvestPlantTags:		harvestDetails.HarvestPlantTags,
			HarvestPlantConsumed:	harvestDetails.HarvestPlantConsumed,
			Status:					harvestDetails.Status}
		harvestDetailAsBytes, _ := json.Marshal(harvestDetail)
		err = stub.PutState(args[0], harvestDetailAsBytes)
		if err != nil {
			return shim.Error("Error while put state to update harvestDetail")
		}
	} else {
		return shim.Error("HarvestId" + args[0] + " doesn't exist. Register the Harvest first.")
	}
	return shim.Success([]byte("HarvestId " + args[0] + " tested successfully."))
}


func (t *SimpleChaincode) viewHarvestDtlsByQALabId(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	harvestIdIterator, err := stub.GetQueryResult("{\"selector\":{\"docType\":\"HarvestDetails\",\"qaLabId\":\""+ args[0] + "\"}}")
	if err != nil {
		return shim.Error(err.Error())
	}

	defer harvestIdIterator.Close()
	// buffer is a JSON array containing QueryResults
	var harvestIdBuffer bytes.Buffer
	harvestIdBuffer.WriteString("[")

	harvestIdAlreadyWritten := false
	for harvestIdIterator.HasNext() {
		queryResponse, err := harvestIdIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if harvestIdAlreadyWritten == true {
			harvestIdBuffer.WriteString(",")
		}
		harvestIdBuffer.WriteString("{\"Key\":")
		harvestIdBuffer.WriteString("\"")
		harvestIdBuffer.WriteString(queryResponse.Key)
		harvestIdBuffer.WriteString("\"")

		harvestIdBuffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		harvestIdBuffer.WriteString(string(queryResponse.Value))
		harvestIdBuffer.WriteString("}")
		harvestIdAlreadyWritten = true
	}
	harvestIdBuffer.WriteString("]")

	return shim.Success(harvestIdBuffer.Bytes())	
}


func (t *SimpleChaincode) viewHarvestDtlsByGrade(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	
	harvestIdIterator, err := stub.GetQueryResult("{\"selector\":{\"docType\":\"HarvestDetails\",\"grade\":\"D\"}}")
	if err != nil {
		return shim.Error(err.Error())
	}

	defer harvestIdIterator.Close()
	// buffer is a JSON array containing QueryResults
	var harvestIdBuffer bytes.Buffer
	harvestIdBuffer.WriteString("[")

	harvestIdAlreadyWritten := false
	for harvestIdIterator.HasNext() {
		queryResponse, err := harvestIdIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if harvestIdAlreadyWritten == true {
			harvestIdBuffer.WriteString(",")
		}
		harvestIdBuffer.WriteString("{\"Key\":")
		harvestIdBuffer.WriteString("\"")
		harvestIdBuffer.WriteString(queryResponse.Key)
		harvestIdBuffer.WriteString("\"")

		harvestIdBuffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		harvestIdBuffer.WriteString(string(queryResponse.Value))
		harvestIdBuffer.WriteString("}")
		harvestIdAlreadyWritten = true
	}
	harvestIdBuffer.WriteString("]")

	return shim.Success(harvestIdBuffer.Bytes())	
}


func (t *SimpleChaincode) destroyHarvest(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	harvestIdIterator, err := stub.GetQueryResult("{\"selector\":{\"docType\":\"HarvestDetails\",\"harvestId\":\"" + args[0] + "\"}}")
	if err != nil {
		return shim.Error(err.Error())
	}

	if harvestIdIterator.HasNext() {
		var err error
		resultQuery, resultError := harvestIdIterator.Next()
		if resultError != nil {
			return shim.Error(resultError.Error())
		}
		var harvestDetails HarvestDetails
		unMarshalError := json.Unmarshal(resultQuery.Value, &harvestDetails)
		if unMarshalError != nil {
			return shim.Error(unMarshalError.Error())
		}
		
		var harvestDetail = HarvestDetails{
			AssetType:				"HarvestDetails",     		
			HarvestID:				harvestDetails.HarvestID,
			HarvestType:			harvestDetails.HarvestType,
			LotCapacity:			harvestDetails.LotCapacity,
			TestingStatus:			harvestDetails.TestingStatus,
			Grade:					harvestDetails.Grade,
			TestingRemarks:			harvestDetails.TestingRemarks,
			CreatedDate:			harvestDetails.CreatedDate,
			DispatchedDate:			harvestDetails.DispatchedDate,
			TestingDate:			harvestDetails.TestingDate,
			QALabID:				harvestDetails.QALabID,
			DestroyerQALabID:		args[3],
			CultivatorID:			harvestDetails.CultivatorID,
			CertificateID:			harvestDetails.CertificateID,
			ExpiryDate:				harvestDetails.ExpiryDate,
			DestroyedDate:			args[2],
			DestroyedRemarks:		args[1],
			HarvestPlantTags:		harvestDetails.HarvestPlantTags,
			HarvestPlantConsumed:	harvestDetails.HarvestPlantConsumed,
			Status:					harvestDetails.Status}
		harvestDetailAsBytes, _ := json.Marshal(harvestDetail)
		err = stub.PutState(args[0], harvestDetailAsBytes)
		if err != nil {
			return shim.Error("Error while put state to update harvestDetail")
		}
	} else {
		return shim.Error("HarvestId" + args[0] + " doesn't exist. Register the Harvest first.")
	}
	return shim.Success([]byte("HarvestId " + args[0] + " destroyed successfully."))
}

func (t *SimpleChaincode) splitHarvest(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	
	harvestIdIterator, err := stub.GetQueryResult("{\"selector\":{\"docType\":\"HarvestDetails\",\"qaLabId\":\"" + args[0] + "\",\"testingStatus\":\"Pass\"}}")
	if err != nil {
		return shim.Error(err.Error())
	}

	defer harvestIdIterator.Close()
	// buffer is a JSON array containing QueryResults
	var harvestIdBuffer bytes.Buffer
	harvestIdBuffer.WriteString("[")

	harvestIdAlreadyWritten := false
	for harvestIdIterator.HasNext() {
		queryResponse, err := harvestIdIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if harvestIdAlreadyWritten == true {
			harvestIdBuffer.WriteString(",")
		}
		harvestIdBuffer.WriteString("{\"Key\":")
		harvestIdBuffer.WriteString("\"")
		harvestIdBuffer.WriteString(queryResponse.Key)
		harvestIdBuffer.WriteString("\"")

		harvestIdBuffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		harvestIdBuffer.WriteString(string(queryResponse.Value))
		harvestIdBuffer.WriteString("}")
		harvestIdAlreadyWritten = true
	}
	harvestIdBuffer.WriteString("]")

	return shim.Success(harvestIdBuffer.Bytes())	
}

func (t *SimpleChaincode) createLot(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}
	var harvestArray[] string
	var lotIdArray[] string
	harvestArrayAsbytes:= []byte(args[0])
	unMarshalHarvestError:= json.Unmarshal(harvestArrayAsbytes, &harvestArray)
	if unMarshalHarvestError != nil {
		return shim.Error(unMarshalHarvestError.Error())
	}
		
	lotIdArrayAsbytes:= []byte(args[1])
	unMarshalLotError:=json.Unmarshal(lotIdArrayAsbytes, &lotIdArray)
	if unMarshalLotError != nil {
		return shim.Error(unMarshalLotError.Error())
	}

	var precise = 0
	
	for i:=0; i<len(harvestArray); i++ {
		var harvestLotArray[] string
		harvestIdIterator, err := stub.GetQueryResult("{\"selector\":{\"docType\":\"HarvestDetails\",\"harvestId\":\"" + harvestArray[i] + "\"}}")
		if err != nil {
			return shim.Error(err.Error())
		}

		if harvestIdIterator.HasNext() {
			var err error
			resultQuery, resultError := harvestIdIterator.Next()
			if resultError != nil {
				return shim.Error(resultError.Error())
			}
			var harvestDetails HarvestDetails
			unMarshalError := json.Unmarshal(resultQuery.Value, &harvestDetails)
			if unMarshalError != nil {
				return shim.Error(unMarshalError.Error())
			}
			limit,err := strconv.Atoi(harvestDetails.LotCapacity)
			if err!= nil {
				fmt.Println("String conversion error")
			}
			if harvestDetails.HarvestType =="Bud" {
				for j:= 0; j< limit; j++ {
				var lotDetails= LotDetails {
					AssetType:		"LotDetails",  
					LotID:			lotIdArray[precise+j],
					LotType:		harvestDetails.HarvestType,
					Grade:			harvestDetails.Grade,
					CreatedDate:	args[2],
					TestingDate:	harvestDetails.TestingDate,
					TestingRemarks:	harvestDetails.TestingRemarks,
					Status:			"R",
					PacketCapacity: "16"}
					lotDetailAsBytes, _ := json.Marshal(lotDetails)
					err = stub.PutState(lotIdArray[precise+j], lotDetailAsBytes)
					if err != nil {
						return shim.Error("Error while put state of Lot Detail for LotId " +lotIdArray[precise+j])
					}
					harvestLotArray= append(harvestLotArray, lotIdArray[precise+j])
				
				}
			

		}	else { 
				for j:= 0; j< limit; j++ {
				var lotDetails= LotDetails {
					AssetType:		"LotDetails",  
					LotID:			lotIdArray[precise+j],
					LotType:		harvestDetails.HarvestType,
					Grade:			harvestDetails.Grade,
					CreatedDate:	args[2],
					TestingDate:	harvestDetails.TestingDate,
					TestingRemarks:	harvestDetails.TestingRemarks,
					Status:			"R",
					PacketCapacity: "8"}
					lotDetailAsBytes, _ := json.Marshal(lotDetails)
					err = stub.PutState(lotIdArray[precise+j], lotDetailAsBytes)
					if err != nil {
						return shim.Error("Error while put state of Lot Detail for LotId " +lotIdArray[precise+j])
					}
					harvestLotArray= append(harvestLotArray, lotIdArray[precise+j])

					}
				}
				
			
			var harvestDetail = HarvestDetails{
				AssetType:				"HarvestDetails",     		
				HarvestID:				harvestDetails.HarvestID,
				HarvestType:			harvestDetails.HarvestType,
				LotCapacity:			harvestDetails.LotCapacity,
				TestingStatus:			harvestDetails.TestingStatus,
				Grade:					harvestDetails.Grade,
				TestingRemarks:			harvestDetails.TestingRemarks,
				CreatedDate:			harvestDetails.CreatedDate,
				DispatchedDate:			harvestDetails.DispatchedDate,
				TestingDate:			harvestDetails.TestingDate,
				QALabID:				harvestDetails.QALabID,
				CultivatorID:			harvestDetails.CultivatorID,
				CertificateID:			harvestDetails.CertificateID,
				ExpiryDate:				harvestDetails.ExpiryDate,
				DestroyedDate:			harvestDetails.DestroyedDate,
				DestroyedRemarks:		harvestDetails.DestroyedRemarks,
				HarvestPlantTags:		harvestDetails.HarvestPlantTags,
				HarvestPlantConsumed:	harvestDetails.HarvestPlantConsumed,
				LotInfo:				harvestLotArray,
				Status:					harvestDetails.Status}
			harvestDetailAsBytes, _ := json.Marshal(harvestDetail)
			err = stub.PutState(harvestArray[i], harvestDetailAsBytes)
			if err != nil {
				return shim.Error("Error while put state to update harvestDetail while creating Lots")
			}
			precise+=limit
		}
		
	}
	return shim.Success([]byte("Lots created successfully."))
}


func (t *SimpleChaincode) viewLotDtls(stub shim.ChaincodeStubInterface) pb.Response {
	

	lotIdIterator, err := stub.GetQueryResult("{\"selector\":{\"docType\":\"LotDetails\"}}")
	if err != nil {
		return shim.Error(err.Error())
	}

	defer lotIdIterator.Close()
	// buffer is a JSON array containing QueryResults
	var lotIdBuffer bytes.Buffer
	lotIdBuffer.WriteString("[")

	lotIdAlreadyWritten := false
	for lotIdIterator.HasNext() {
		queryResponse, err := lotIdIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if lotIdAlreadyWritten == true {
			lotIdBuffer.WriteString(",")
		}
		lotIdBuffer.WriteString("{\"Key\":")
		lotIdBuffer.WriteString("\"")
		lotIdBuffer.WriteString(queryResponse.Key)
		lotIdBuffer.WriteString("\"")

		lotIdBuffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		lotIdBuffer.WriteString(string(queryResponse.Value))
		lotIdBuffer.WriteString("}")
		lotIdAlreadyWritten = true
	}
	lotIdBuffer.WriteString("]")

	return shim.Success(lotIdBuffer.Bytes())	
}

func (t *SimpleChaincode) dispatchLot(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}
	var lotIdArray[] string 
	lotIdArrayAsbytes:= []byte(args[0])
	unMarshalError:= json.Unmarshal(lotIdArrayAsbytes, &lotIdArray)
	if unMarshalError!= nil {
		return shim.Error(unMarshalError.Error())
	}
	for i:=0; i<len(lotIdArray); i++ {
		lotIdIterator, err := stub.GetQueryResult("{\"selector\":{\"docType\":\"LotDetails\",\"lotId\":\"" + lotIdArray[i] + "\"}}")
		if err != nil {
			return shim.Error(err.Error())
		}
		
		if lotIdIterator.HasNext() {
		var err error
		resultQuery, resultError := lotIdIterator.Next()
		if resultError != nil {
			return shim.Error(resultError.Error())
		}
		var lotDetails LotDetails
		unMarshalLotError := json.Unmarshal(resultQuery.Value, &lotDetails)
		if unMarshalLotError != nil {
			return shim.Error(unMarshalLotError.Error())
		}
		
		var lotDetail = LotDetails{
			AssetType:			"LotDetails",     		
			LotID:				lotDetails.LotID,
			LotType:			lotDetails.LotType,
			Grade:				lotDetails.Grade,
			QALabID:			args[1],
			CreatedDate:		lotDetails.CreatedDate,
			DispatchedDate:		args[4],
			PackagerID:			args[2],
			LotBatchID:			args[3],
			TestingDate:		lotDetails.TestingDate,
			TestingRemarks:		lotDetails.TestingRemarks,
			Status:				"D",
			PacketCapacity:		lotDetails.PacketCapacity}
		lotDetailAsBytes, _ := json.Marshal(lotDetail)
		err = stub.PutState(lotIdArray[i], lotDetailAsBytes)
		if err != nil {
			return shim.Error("Error while put state to dispatch lotDetail")
		}
	} else {
		return shim.Error("LotId " + lotIdArray[i] + " doesn't exist.Create the Lot first.")
	}

	}
	return shim.Success([]byte("Lots has been dispatched against BatchId "+args[3] ))
}

// func (t *SimpleChaincode) viewLotDtlsByPackagerId(stub shim.ChaincodeStubInterface, args []string) pb.Response {
// 	//change
// 	if len(args) != 1 {
// 		return shim.Error("Incorrect number of arguments. Expecting 1")
// 	}
// 	lotIdIterator, err := stub.GetQueryResult("{\"selector\":{\"docType\":\"LotDetails\",\"packagerId\":\"" + args[0] + "\"}}")
// 	if err != nil {
// 		return shim.Error(err.Error())
// 	}

// 	defer lotIdIterator.Close()
// 	// buffer is a JSON array containing QueryResults
// 	var lotIdBuffer bytes.Buffer
// 	lotIdBuffer.WriteString("[")

// 	lotIdAlreadyWritten := false
// 	for lotIdIterator.HasNext() {
// 		queryResponse, err := lotIdIterator.Next()
// 		if err != nil {
// 			return shim.Error(err.Error())
// 		}
// 		// Add a comma before array members, suppress it for the first array member
// 		if lotIdAlreadyWritten == true {
// 			lotIdBuffer.WriteString(",")
// 		}
// 		lotIdBuffer.WriteString("{\"Key\":")
// 		lotIdBuffer.WriteString("\"")
// 		lotIdBuffer.WriteString(queryResponse.Key)
// 		lotIdBuffer.WriteString("\"")

// 		lotIdBuffer.WriteString(", \"Record\":")
// 		// Record is a JSON object, so we write as-is
// 		lotIdBuffer.WriteString(string(queryResponse.Value))
// 		lotIdBuffer.WriteString("}")
// 		lotIdAlreadyWritten = true
// 	}
// 	lotIdBuffer.WriteString("]")

// 	return shim.Success(lotIdBuffer.Bytes())	
// }

func (t *SimpleChaincode) viewLotDtlsByPackagerId(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//change
	if len(args) == 1 {

	lotIdIterator, err := stub.GetQueryResult("{\"selector\":{\"docType\":\"LotDetails\",\"packagerId\":\"" + args[0] + "\"}}")
	if err != nil {
		return shim.Error(err.Error())
	}

	defer lotIdIterator.Close()
	// buffer is a JSON array containing QueryResults
	var lotIdBuffer bytes.Buffer
	lotIdBuffer.WriteString("[")

	lotIdAlreadyWritten := false
	for lotIdIterator.HasNext() {
		queryResponse, err := lotIdIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if lotIdAlreadyWritten == true {
			lotIdBuffer.WriteString(",")
		}
		lotIdBuffer.WriteString("{\"Key\":")
		lotIdBuffer.WriteString("\"")
		lotIdBuffer.WriteString(queryResponse.Key)
		lotIdBuffer.WriteString("\"")

		lotIdBuffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		lotIdBuffer.WriteString(string(queryResponse.Value))
		lotIdBuffer.WriteString("}")
		lotIdAlreadyWritten = true
	}
	lotIdBuffer.WriteString("]")

	return shim.Success(lotIdBuffer.Bytes())	


	} else if len(args) == 2 { 
	lotIdIterator, err := stub.GetQueryResult("{\"selector\":{\"docType\":\"LotDetails\",\"packagerId\":\"" + args[0] + "\"}}")
	if err != nil {
		return shim.Error(err.Error())
	}

	defer lotIdIterator.Close()
	// buffer is a JSON array containing QueryResults
	var lotIdBuffer bytes.Buffer
	lotIdBuffer.WriteString("[")

	lotIdAlreadyWritten := false
	for lotIdIterator.HasNext() {
		queryResponse, err := lotIdIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if lotIdAlreadyWritten == true {
			lotIdBuffer.WriteString(",")
		}
		lotIdBuffer.WriteString("{\"Key\":")
		lotIdBuffer.WriteString("\"")
		lotIdBuffer.WriteString(queryResponse.Key)
		lotIdBuffer.WriteString("\"")

		lotIdBuffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		lotIdBuffer.WriteString(string(queryResponse.Value))
		lotIdBuffer.WriteString("}")
		lotIdAlreadyWritten = true
	}
	lotIdBuffer.WriteString("]")

	// return shim.Success(lotIdBuffer.Bytes())
	packetIdIterator, err := stub.GetQueryResult("{\"selector\":{\"docType\":\"PacketDetails\",\"packagerId\":\""+ args[0] + "\"}}")
	if err != nil {
		return shim.Error(err.Error())
	}
	defer packetIdIterator.Close()

	var packetIdBuffer bytes.Buffer
	packetIdBuffer.WriteString("[")
	//start composing the response to be sent to node
	packetIdAlreadyWritten := false
	for packetIdIterator.HasNext() {
		queryResponse, err := packetIdIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if packetIdAlreadyWritten == true {
			packetIdBuffer.WriteString(",")
		}
		packetIdBuffer.WriteString("{\"Key\":")
		packetIdBuffer.WriteString("\"")
		packetIdBuffer.WriteString(queryResponse.Key)
		packetIdBuffer.WriteString("\"")

		packetIdBuffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		packetIdBuffer.WriteString(string(queryResponse.Value))
		packetIdBuffer.WriteString("}")
		packetIdAlreadyWritten = true
	}
	packetIdBuffer.WriteString("]")

	var res = PacketLotResponse{LotAsset: lotIdBuffer.String(), PacketAsset: packetIdBuffer.String()}
	finalResponse, marshalerr := json.Marshal(res)
	if marshalerr != nil {
		shim.Error(marshalerr.Error())
	}
	return shim.Success(finalResponse)

} else {
	return shim.Error("Incorrect number of arguments. Expecting 1 or 2")
}

}


func (t *SimpleChaincode) createPacket(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}
	var lotArray[] string
	var packetIdArray[] string

	lotArrayAsBytes:=[]byte(args[0])
	unMarshalLotError:=json.Unmarshal(lotArrayAsBytes, &lotArray)
	if unMarshalLotError != nil {
		return shim.Error(unMarshalLotError.Error())
	}

	packetIdArrayAsBytes:=[]byte(args[1])
	unMarshallPacketError:=json.Unmarshal(packetIdArrayAsBytes, &packetIdArray)
	if unMarshallPacketError != nil {
		return shim.Error(unMarshallPacketError.Error())

	}
	//we have extracted the lots selected to make packets
	//and packetIds generated in node,now we'll select the lot according to the lot id we extracted

	var precise=0

	for i:=0; i<len(lotArray); i++ {
		var lotPacketArray[] string
		// to store the packetIds generated from that particular lot
		lotIdIterator,err:=stub.GetQueryResult("{\"selector\":{\"docType\":\"LotDetails\",\"lotId\":\"" + lotArray[i] + "\"}}")
		if err != nil {
			return shim.Error(err.Error())
		}

		if lotIdIterator.HasNext() {
			var err error
			resultQuery,resultError := lotIdIterator.Next()
			//resultQuery stores a lot asset with a particular lotId selected from the lotArray
			if resultError != nil {
				return shim.Error(resultError.Error())
			}
			var lotDetails LotDetails
			unMarshalError:= json.Unmarshal(resultQuery.Value, &lotDetails)
			if unMarshalError != nil {
				return shim.Error(unMarshalError.Error())
			}
			limit,err := strconv.Atoi(lotDetails.PacketCapacity)
			if err != nil {
				fmt.Println("String Conversion Error")
			}

			for j :=0;j<limit; j++ {
				var packetDetails= PacketDetails {
					AssetType:			"PacketDetails",
					PacketCreationDate:	args[2],
					PacketID:			packetIdArray[precise+j],
					LotID:				lotDetails.LotID,
					Type:				lotDetails.LotType,
					Grade:				lotDetails.Grade,
					PacketBatchID:		args[3],
					PackagerID:			lotDetails.PackagerID,
					Dispatched:			"N"}

				
				packetDetailsAsBytes, _:= json.Marshal(packetDetails)
				err = stub.PutState(packetIdArray[precise+j], packetDetailsAsBytes)
				if err != nil {
					return shim.Error("Error while put state of Packet Detail for PacketId " +packetIdArray[precise+j])
				}
				lotPacketArray=append(lotPacketArray, packetIdArray[precise+j])
			
			}
			var lotDetail = LotDetails{
				AssetType:				"LotDetails",			
				LotID:					lotDetails.LotID,			
				LotType:				lotDetails.LotType,			
				Grade:					lotDetails.Grade,			
				QALabID:				lotDetails.QALabID,			
				CreatedDate:			lotDetails.CreatedDate,			
				DispatchedDate:			lotDetails.DispatchedDate,			
				PackagerID:				lotDetails.PackagerID,			
				LotBatchID:				lotDetails.LotBatchID,			
				TestingDate:			lotDetails.TestingDate,			
				TestingRemarks:			lotDetails.TestingRemarks,			
				Status:					lotDetails.Status,		
				PacketCapacity:			lotDetails.PacketCapacity,			
				PacketInfo:				lotPacketArray}
			
			lotDetailAsBytes, _ := json.Marshal(lotDetail)
			err = stub.PutState(lotArray[i], lotDetailAsBytes)
			if err != nil {
				return shim.Error("Error while put state to update harvestDetail while creating Lots")
			}
			precise+=limit
		}
	}
	return shim.Success([]byte("Packets created successfully."))
}

func (t *SimpleChaincode) viewPacketDtlsByLotId(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	packetIdIterator, err := stub.GetQueryResult("{\"selector\":{\"docType\":\"PacketDetails\",\"lotId\":\""+ args[0] + "\"}}")
	if err != nil {
		return shim.Error(err.Error())
	}
	defer packetIdIterator.Close()

	var packetIdBuffer bytes.Buffer
	packetIdBuffer.WriteString("[")
	//start composing the response to be sent to node
	packetIdAlreadyWritten := false
	for packetIdIterator.HasNext() {
		queryResponse, err := packetIdIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if packetIdAlreadyWritten == true {
			packetIdBuffer.WriteString(",")
		}
		packetIdBuffer.WriteString("{\"Key\":")
		packetIdBuffer.WriteString("\"")
		packetIdBuffer.WriteString(queryResponse.Key)
		packetIdBuffer.WriteString("\"")

		packetIdBuffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		packetIdBuffer.WriteString(string(queryResponse.Value))
		packetIdBuffer.WriteString("}")
		packetIdAlreadyWritten = true
	}
	packetIdBuffer.WriteString("]")

	return shim.Success(packetIdBuffer.Bytes())

}

func (t *SimpleChaincode) viewPurchaseOrderByPackagerId(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	purchaseOrderIdIterator, err := stub.GetQueryResult("{\"selector\":{\"docType\":\"PurchaseOrderDetails\",\"packagerId\":\""+ args[0] + "\"}}")
	if err != nil {
		return shim.Error(err.Error())
	}
	defer purchaseOrderIdIterator.Close()

	var purchaseOrderIdBuffer bytes.Buffer
	purchaseOrderIdBuffer.WriteString("[")

	purchaseOrderIdAlreadyWritten := false
	for purchaseOrderIdIterator.HasNext() {
		queryResponse, err := purchaseOrderIdIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if purchaseOrderIdAlreadyWritten == true {
			purchaseOrderIdBuffer.WriteString(",")
		}
		purchaseOrderIdBuffer.WriteString("{\"Key\":")
		purchaseOrderIdBuffer.WriteString("\"")
		purchaseOrderIdBuffer.WriteString(queryResponse.Key)
		purchaseOrderIdBuffer.WriteString("\"")

		purchaseOrderIdBuffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		purchaseOrderIdBuffer.WriteString(string(queryResponse.Value))
		purchaseOrderIdBuffer.WriteString("}")
		purchaseOrderIdAlreadyWritten = true
	}
	purchaseOrderIdBuffer.WriteString("]")

	//subsequent steps to send packet assets with Dispatched field "N"
	packetIdIterator, err := stub.GetQueryResult("{\"selector\":{\"docType\":\"PacketDetails\",\"packagerId\":\"" + args[0] + "\"}}")
	if err != nil {
		return shim.Error(err.Error())
	}

	defer packetIdIterator.Close()
	// buffer is a JSON array containing QueryResults
	var packetIdBuffer bytes.Buffer
	packetIdBuffer.WriteString("[")

	packetIdAlreadyWritten := false
	for packetIdIterator.HasNext() {
		queryResponse, err := packetIdIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if packetIdAlreadyWritten == true {
			packetIdBuffer.WriteString(",")
		}
		packetIdBuffer.WriteString("{\"Key\":")
		packetIdBuffer.WriteString("\"")
		packetIdBuffer.WriteString(queryResponse.Key)
		packetIdBuffer.WriteString("\"")

		packetIdBuffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		packetIdBuffer.WriteString(string(queryResponse.Value))
		packetIdBuffer.WriteString("}")
		packetIdAlreadyWritten = true
	}
	packetIdBuffer.WriteString("]")

	var res = PacketResponse{PurchaseOrderAsset: purchaseOrderIdBuffer.String(), PacketAsset: packetIdBuffer.String()}
	finalResponse, marshalerr := json.Marshal(res)
	if marshalerr != nil {
		shim.Error(marshalerr.Error())
	}
	return shim.Success(finalResponse)

}

func (t *SimpleChaincode) dispatchPacket(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}
	var count int
	count=0
	var packetIdArray[] string 
	packetIterator, err := stub.GetQueryResult("{\"selector\":{\"docType\":\"PacketDetails\",\"packagerId\":\"" + args[0] + "\",\"dispatched\":\"N\",\"type\":\""+ args[3] +"\",\"grade\":\""+ args[4] +"\"}}")
	if err != nil {
		return shim.Error(err.Error())
	}
	for packetIterator.HasNext() {
		count++
		logger.Info("count value---------------------",count)
		resultQuery, resultError := packetIterator.Next()
		if resultError != nil {
			return shim.Error(resultError.Error())
		}
		var packetDetails PacketDetails
		unMarshalPacketError := json.Unmarshal(resultQuery.Value, &packetDetails)
		if unMarshalPacketError != nil {
			return shim.Error(unMarshalPacketError.Error())
		}
		logger.Info(" packetDetails---------------------",packetDetails)
		logger.Info(" packetDetails.PacketID---------------------",packetDetails.PacketID)
		packetIdArray=append(packetIdArray, packetDetails.PacketID)
		logger.Info(" packetIdArray---------------------",packetIdArray)

	}
	
	var purchaseOrderIdArray[] string 
	purchaseOrderIdArrayAsbytes:= []byte(args[1])
	unMarshalError:= json.Unmarshal(purchaseOrderIdArrayAsbytes, &purchaseOrderIdArray)
	logger.Info(" purchaseOrderIdArray---------------------",purchaseOrderIdArray)
	if unMarshalError!= nil {
		return shim.Error(unMarshalError.Error())
	}

	var precise = 0

	for i:=0; i<len(purchaseOrderIdArray); i++ {
		var purchaseOrderPacketArray[] string
		//to store all the packets to PacketsInfo field
		purchaseOrderIdIterator, err := stub.GetQueryResult("{\"selector\":{\"docType\":\"PurchaseOrderDetails\",\"purchaseOrderId\":\"" + purchaseOrderIdArray[i] + "\"}}")
		if err != nil {
			return shim.Error(err.Error())
		}

		if purchaseOrderIdIterator.HasNext() {
			var err error
			resultQuery, resultError := purchaseOrderIdIterator.Next()
		if resultError != nil {
			return shim.Error(resultError.Error())
		}
		var purchaseOrderDetails PurchaseOrderDetails
		unMarshalLotError := json.Unmarshal(resultQuery.Value, &purchaseOrderDetails)
		if unMarshalLotError != nil {
			return shim.Error(unMarshalLotError.Error())
		}

		limit,err := strconv.Atoi(purchaseOrderDetails.Quantity)
		logger.Info(" limit---------------------",limit)
			if err!= nil {
				fmt.Println("String conversion error")
			}
		

		for j:= 0; j< limit; j++ {

		packetIdIterator, err := stub.GetQueryResult("{\"selector\":{\"docType\":\"PacketDetails\",\"packetId\":\""+ packetIdArray[precise+j] + "\"}}")
			if err != nil {
			return shim.Error(err.Error())
			}
			if packetIdIterator.HasNext() {
				var err error
				resultQuery, resultError := packetIdIterator.Next()
				if resultError != nil {
					return shim.Error(resultError.Error())
				}
				var packetDetails PacketDetails
				unMarshalError := json.Unmarshal(resultQuery.Value, &packetDetails)
				if unMarshalError != nil {
					return shim.Error(unMarshalError.Error())
				}
				logger.Info(" PacketId current---------------------",packetDetails.PacketID)
				var packetDetail= PacketDetails {
				AssetType:			"PacketDetails",
				PacketCreationDate:	packetDetails.PacketCreationDate,
				// PacketID:		packetIdArray[precise+j],
				PacketID:			packetDetails.PacketID,
				LotID:				packetDetails.LotID,
				Type:				packetDetails.Type,
				Grade:				packetDetails.Grade,
				PacketBatchID:		packetDetails.PacketBatchID,
				PackagerID:			packetDetails.PackagerID,
				Dispatched:			"Y"}
				packetDetailsAsBytes, _ := json.Marshal(packetDetail)
				
				err = stub.PutState(packetIdArray[precise+j], packetDetailsAsBytes)
				if err != nil {
					return shim.Error("Error while put state of Packet Detail for packetId " +packetIdArray[precise+j])
				}
				logger.Info(" Packet Put state successful---------------------")
				purchaseOrderPacketArray= append(purchaseOrderPacketArray,packetIdArray[precise+j])
				logger.Info(" purchaseOrderPacketArray---------------------",purchaseOrderPacketArray)
				//this array empties every time the i loop starts afresh,so no prob of having prev packetIds in purchase asset

			}
		}
		
		var purchaseOrderDetail = PurchaseOrderDetails {
			AssetType:			"PurchaseOrderDetails",     		
			PurchaseOrderID:	purchaseOrderDetails.PurchaseOrderID,
			Type:				purchaseOrderDetails.Type,
			Grade:				purchaseOrderDetails.Grade,
			PackagerID:			purchaseOrderDetails.PackagerID,
			PacketInfo:			purchaseOrderPacketArray,
			Quantity:			purchaseOrderDetails.Quantity,
			DispensaryID:		purchaseOrderDetails.DispensaryID,
			OrderDate:			purchaseOrderDetails.OrderDate,
			DispatchedDate:		args[2]}
			
		
		purchaseOrderDetailsBytes, _ := json.Marshal(purchaseOrderDetail)
			err = stub.PutState(purchaseOrderIdArray[i], purchaseOrderDetailsBytes)
			if err != nil {
				return shim.Error("Error while put state to update purchaseOrderDetails while creating Packets")
			}
			precise+=limit
		}
	}	
	return shim.Success([]byte("Packets Dispatched successfully."))


}


func (t *SimpleChaincode) viewDispatchedPurchaseOrder(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	
	if len(args) !=  1{
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	purchaseOrderIterator, err := stub.GetQueryResult("{\"selector\":{\"docType\":\"PurchaseOrderDetails\",\"packagerId\":\"" + args[0] + "\"}}")
	if err != nil {
		return shim.Error(err.Error())
	}

	defer purchaseOrderIterator.Close()
	// buffer is a JSON array containing QueryResults
	var purchaseOrderBuffer bytes.Buffer
	purchaseOrderBuffer.WriteString("[")

	purchaseOrderAlreadyWritten := false
	for purchaseOrderIterator.HasNext() {
		queryResponse, err := purchaseOrderIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if purchaseOrderAlreadyWritten == true {
			purchaseOrderBuffer.WriteString(",")
		}
		purchaseOrderBuffer.WriteString("{\"Key\":")
		purchaseOrderBuffer.WriteString("\"")
		purchaseOrderBuffer.WriteString(queryResponse.Key)
		purchaseOrderBuffer.WriteString("\"")

		purchaseOrderBuffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		purchaseOrderBuffer.WriteString(string(queryResponse.Value))
		purchaseOrderBuffer.WriteString("}")
		purchaseOrderAlreadyWritten = true
	}
	purchaseOrderBuffer.WriteString("]")

	// return shim.Success(purchaseOrderBuffer.Bytes())	
	packetIterator, err := stub.GetQueryResult("{\"selector\":{\"docType\":\"PacketDetails\",\"packagerId\":\"" + args[0] + "\",\"dispatched\":\"Y\"}}")
	if err != nil {							
		return shim.Error(err.Error())
	}
	defer packetIterator.Close()
	// buffer is a JSON array containing QueryResults
	var packetBuffer bytes.Buffer
	packetBuffer.WriteString("[")

	packetAlreadyWritten := false
	for packetIterator.HasNext() {
		queryResponse, err := packetIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		if packetAlreadyWritten == true {
			packetBuffer.WriteString(",")
		}
		packetBuffer.WriteString("{\"Key\":")
		packetBuffer.WriteString("\"")
		packetBuffer.WriteString(queryResponse.Key)
		packetBuffer.WriteString("\"")

		packetBuffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		packetBuffer.WriteString(string(queryResponse.Value))
		packetBuffer.WriteString("}")
		packetAlreadyWritten = true
	}
	packetBuffer.WriteString("]")

	var res = PurchaseResponse{PurchaseOrderAsset: purchaseOrderBuffer.String(), PacketAsset: packetBuffer.String()}
	finalResponse, marshalerr := json.Marshal(res)
	if marshalerr != nil {
		shim.Error(marshalerr.Error())
	}
	return shim.Success(finalResponse)
}



func (t *SimpleChaincode) createPurchaseOrder(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 8 {
		return shim.Error("Incorrect number of arguments. Expecting 8")
	}

	var err error
	var purchaseOrderdetail= PurchaseOrderDetails {
			AssetType:			"PurchaseOrderDetails",     		
			PurchaseOrderID:	args[7],
			Type:				args[0],
			Grade:				args[1],
			Quantity:			args[2],
			OrderDescription:	args[3],
			PackagerID:			args[4],
			DispensaryID:		args[5],
			OrderDate:			args[6]}
			purchaseOrderDetailAsBytes, _ := json.Marshal(purchaseOrderdetail)
			err = stub.PutState(args[7], purchaseOrderDetailAsBytes)
			if err != nil {
				return shim.Error("Error while put state to create purchaseOrderdetail")
			}
			return shim.Success([]byte("PurchaseOrderId " + args[7] + " created Successfully."))
	}

func (t *SimpleChaincode) viewPurchaseOrderByDispensaryId(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) !=  1{
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	purchaseOrderIterator, err := stub.GetQueryResult("{\"selector\":{\"docType\":\"PurchaseOrderDetails\",\"dispensaryId\":\"" + args[0] + "\"}}")
	if err != nil {
		return shim.Error(err.Error())
	}

	defer purchaseOrderIterator.Close()
	// buffer is a JSON array containing QueryResults
	var purchaseOrderBuffer bytes.Buffer
	purchaseOrderBuffer.WriteString("[")

	purchaseOrderAlreadyWritten := false
	for purchaseOrderIterator.HasNext() {
		queryResponse, err := purchaseOrderIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if purchaseOrderAlreadyWritten == true {
			purchaseOrderBuffer.WriteString(",")
		}
		purchaseOrderBuffer.WriteString("{\"Key\":")
		purchaseOrderBuffer.WriteString("\"")
		purchaseOrderBuffer.WriteString(queryResponse.Key)
		purchaseOrderBuffer.WriteString("\"")

		purchaseOrderBuffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		purchaseOrderBuffer.WriteString(string(queryResponse.Value))
		purchaseOrderBuffer.WriteString("}")
		purchaseOrderAlreadyWritten = true
	}
	purchaseOrderBuffer.WriteString("]")

	return shim.Success(purchaseOrderBuffer.Bytes())	
}

func (t *SimpleChaincode) viewPacketByDispensaryId(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) !=  1{
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	var packetIdArray[] string
	var packetIds[] string
	// var packetIdBuffer bytes.Buffer
	
	purchaseOrderIterator, err := stub.GetQueryResult("{\"selector\":{\"docType\":\"PurchaseOrderDetails\",\"dispensaryId\":\"" + args[0] + "\"}}")
		if err != nil {								
			return shim.Error(err.Error())
		}
		for purchaseOrderIterator.HasNext() {
			//var err error
		resultQuery, resultError := purchaseOrderIterator.Next()
		if resultError != nil {
			return shim.Error(resultError.Error())
		}
		var purchaseOrderDetails PurchaseOrderDetails
		unMarshalLotError := json.Unmarshal(resultQuery.Value, &purchaseOrderDetails)
		if unMarshalLotError != nil {
			return shim.Error(unMarshalLotError.Error())
		}
		packetIdArray=purchaseOrderDetails.PacketInfo
		for j:=0;j<len(packetIdArray);j++ {
			packetIds=append(packetIds,packetIdArray[j])
			//we get all the packetsIds Array from all of the purchaseOrders

		}

	}
		var packetIdBuffer bytes.Buffer

		for i:=0;i<len(packetIds);i++ {

			packetIdIterator, err := stub.GetQueryResult("{\"selector\":{\"docType\":\"PacketDetails\",\"packetId\":\"" + packetIds[i] + "\"}}")
			if err != nil {
			return shim.Error(err.Error())
			}
			defer packetIdIterator.Close()
			
			packetIdBuffer.WriteString("[")

			packetIdAlreadyWritten := false
			for packetIdIterator.HasNext() {
			queryResponse, err := packetIdIterator.Next()
			if err != nil {
				return shim.Error(err.Error())
			}
			if packetIdAlreadyWritten == true {
				packetIdBuffer.WriteString(",")
			}
			packetIdBuffer.WriteString("{\"Key\":")
			packetIdBuffer.WriteString("\"")
			packetIdBuffer.WriteString(queryResponse.Key)
			packetIdBuffer.WriteString("\"")
	
			packetIdBuffer.WriteString(", \"Record\":")
			// Record is a JSON object, so we write as-is
			packetIdBuffer.WriteString(string(queryResponse.Value))
			packetIdBuffer.WriteString("}")
			packetIdAlreadyWritten = true
		}
		packetIdBuffer.WriteString("]")
	}
	//return shim.Success(packetIdBuffer.Bytes())

	patientIterator, err := stub.GetQueryResult("{\"selector\":{\"docType\":\"PatientdtByDisPensary\",\"dispensaryId\":\"" + args[0] + "\"}}")
	if err != nil {								
		return shim.Error(err.Error())
	}

	defer patientIterator.Close()
	// buffer is a JSON array containing QueryResults
	var patientBuffer bytes.Buffer
	patientBuffer.WriteString("[")

	patientAlreadyWritten := false
	for patientIterator.HasNext() {
		queryResponse, err := patientIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if patientAlreadyWritten == true {
			patientBuffer.WriteString(",")
		}
		patientBuffer.WriteString("{\"Key\":")
		patientBuffer.WriteString("\"")
		patientBuffer.WriteString(queryResponse.Key)
		patientBuffer.WriteString("\"")

		patientBuffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		patientBuffer.WriteString(string(queryResponse.Value))
		patientBuffer.WriteString("}")
		patientAlreadyWritten = true
	}
	patientBuffer.WriteString("]")

	logger.Info("patientBuffer---------------------",patientBuffer)
	logger.Info("patientBuffer.String()---------------------",patientBuffer.String())


	var res = PatientResponse{PacketAsset: packetIdBuffer.String(),PatientAsset: patientBuffer.String()}
	finalResponse, marshalerr := json.Marshal(res)
	if marshalerr != nil {
		shim.Error(marshalerr.Error())
	}
	return shim.Success(finalResponse)

}


func (t *SimpleChaincode) issuePacket(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 8 {
		return shim.Error("Incorrect number of arguments. Expecting 8")
	}
	var packetIdArray[] string
	packetIdArrayAsBytes:=[]byte(args[0])
	unMarshallPacketError:=json.Unmarshal(packetIdArrayAsBytes, &packetIdArray)
	if unMarshallPacketError != nil {
		return shim.Error(unMarshallPacketError.Error())

	}
	for i:=0; i<len(packetIdArray); i++ {
		packetIdIterator,err:=stub.GetQueryResult("{\"selector\":{\"docType\":\"PacketDetails\",\"packetId\":\"" + packetIdArray[i] + "\"}}")
		if err != nil {
			return shim.Error(err.Error())
		}

		if packetIdIterator.HasNext() {
			var err error
			resultQuery,resultError := packetIdIterator.Next()
			//resultQuery stores a lot asset with a particular lotId selected from the lotArray
			if resultError != nil {
				return shim.Error(resultError.Error())
			}
			var packetDetails PacketDetails
			unMarshalError:= json.Unmarshal(resultQuery.Value, &packetDetails)
			if unMarshalError != nil {
				return shim.Error(unMarshalError.Error())
			}
			var packetDetail= PacketDetails {
				AssetType:			"PacketDetails",
				PacketCreationDate:	packetDetails.PacketCreationDate,
				PacketID:			packetDetails.PacketID,
				LotID:				packetDetails.LotID,
				Type:				packetDetails.Type,
				Grade:				packetDetails.Grade,
				PacketBatchID:		packetDetails.PacketBatchID,
				PackagerID:			packetDetails.PackagerID,
				PatientID: 			args[1],
				Prescription:		args[2],
				IssueDate:			args[7],
				Dispatched:			"Y"}

				packetDetailsAsBytes, _ := json.Marshal(packetDetail)
				err = stub.PutState(packetDetails.PacketID, packetDetailsAsBytes)
				if err != nil {
					return shim.Error("Error while put state of Packet Detail for packetId " +packetDetails.PacketID)
				}

			}

		}
		//return shim.Success([]byte("Packets Issued successfully."))
		//return shim.Success([]byte("Lots has been dispatched against BatchId "+args[3] ))
		patientIterator,err:=stub.GetQueryResult("{\"selector\":{\"docType\":\"PatientdtByDisPensary\",\"dispensaryId\":\"" + args[6] + "\"}}")
		if err != nil {
			return shim.Error(err.Error())
		}
		if patientIterator.HasNext() {
			var err error
				resultQuery, resultError := patientIterator.Next()
				if resultError != nil {
					return shim.Error(resultError.Error())
				}
				var patientDt PatientdtByDisPensary
				unMarshalError := json.Unmarshal(resultQuery.Value, &patientDt)
				if unMarshalError != nil {
					return shim.Error(unMarshalError.Error())
				}
			var patientDetail= PatientDetails {
				AssetType:			"PatientDetails",
				PatientID:			args[1],
				Type:				args[5],
				Grade:				args[4],
				PhoneNumber:		args[3],
				Prescription:		args[2],
				PacketInfo:			packetIdArray,
				DispensaryID:		args[6]}
				logger.Info("patientDetail---------------------",patientDetail)
			patientDt.PatientDetails=append(patientDt.PatientDetails,patientDetail)
			var patientDetailsByDis=PatientdtByDisPensary {
				AssetType:			"PatientdtByDisPensary",
				DispensaryID:		args[6],
				PatientDetails:      patientDt.PatientDetails}
				patientAsBytes, _ := json.Marshal(patientDetailsByDis)
				logger.Info("patientAsBytes---------------------",patientAsBytes)
				err = stub.PutState(args[6], patientAsBytes)
				if err != nil {
					return shim.Error("Error while put state of patient Detail for patientId " +args[1])
				}
		} else   {
			logger.Info("inside else---------------------")
			var temp []PatientDetails
			var patientDetail= PatientDetails {
				AssetType:			"PatientDetails",
				PatientID:			args[1],
				Type:				args[5],
				Grade:				args[4],
				PhoneNumber:		args[3],
				Prescription:		args[2],
				PacketInfo:			packetIdArray,
				DispensaryID:		args[6]}
				logger.Info("patientDetail---------------------",patientDetail)
			temp=append(temp,patientDetail)
			var patientDetailsByDis=PatientdtByDisPensary {
				AssetType:			"PatientdtByDisPensary",
				DispensaryID:		args[6],
				PatientDetails:      temp}
				patientAsBytes, _ := json.Marshal(patientDetailsByDis)
				logger.Info("patientAsBytes---------------------",patientAsBytes)
				err = stub.PutState(args[6], patientAsBytes)
				if err != nil {
					return shim.Error("Error while put state of patient Detail for patientId " +args[1])
				}
			}
			patient,err:=stub.GetQueryResult("{\"selector\":{\"docType\":\"Patient\",\"patientId\":\"" + args[1] + "\"}}")
			if err != nil {
				return shim.Error(err.Error())
			}
			result,Error := patient.Next()
			var patientDt Patient
			if Error != nil{
				return shim.Error(Error.Error())
			}
			unMarshalError:= json.Unmarshal(result.Value, &patientDt)
			if unMarshalError != nil {
				return shim.Error(unMarshalError.Error())
			}
			var patientPres []PatientPrescription
			for _,temp := range patientDt.Patient_Prescription{
				if unMarshalError != nil {
					return shim.Error(unMarshalError.Error())
				}
				var prescrip=temp
				if temp.Status=="Pending" {
				 prescrip=PatientPrescription{
						AssetType:		"patient_Prescription",
						Date:			temp.Date,
						Doses:			temp.Doses,
						Type:			temp.Type,
						Grade:			temp.Grade,
						Status:			"Recieved",
						DoctorID:		temp.DoctorID,
						Prescription:	temp.Prescription}
				}
				patientPres=append(patientPres,prescrip)
			}
			patientDt.Patient_Prescription=patientPres
			patientBytes,_:=json.Marshal(patientDt)
			err = stub.PutState(args[1], patientBytes)
			if err != nil {
				return shim.Error("Error while put state of patient for patientId " +args[1])
			}
			return shim.Success([]byte("Packet issued successfully to patient id " + args[1]))
		}

		func (t *SimpleChaincode) viewPacketDtlsByPacketId(stub shim.ChaincodeStubInterface, args []string) pb.Response {
			if len(args) != 1 {
				return shim.Error("Incorrect number of arguments. Expecting 1")
			}
			packetIdIterator, err := stub.GetQueryResult("{\"selector\":{\"docType\":\"PacketDetails\",\"packetId\":\"" + args[0] + "\"}}")
			if err != nil {
				return shim.Error(err.Error())
			}
		
			defer packetIdIterator.Close()
			// buffer is a JSON array containing QueryResults
			var packetIdBuffer bytes.Buffer
			packetIdBuffer.WriteString("[")
		
			packetIdAlreadyWritten := false
			for packetIdIterator.HasNext() {
				queryResponse, err := packetIdIterator.Next()
				if err != nil {
					return shim.Error(err.Error())
				}
				// Add a comma before array members, suppress it for the first array member
				if packetIdAlreadyWritten == true {
					packetIdBuffer.WriteString(",")
				}
				packetIdBuffer.WriteString("{\"Key\":")
				packetIdBuffer.WriteString("\"")
				packetIdBuffer.WriteString(queryResponse.Key)
				packetIdBuffer.WriteString("\"")
		
				packetIdBuffer.WriteString(", \"Record\":")
				// Record is a JSON object, so we write as-is
				packetIdBuffer.WriteString(string(queryResponse.Value))
				packetIdBuffer.WriteString("}")
				packetIdAlreadyWritten = true
			}
			packetIdBuffer.WriteString("]")
		
			return shim.Success(packetIdBuffer.Bytes())	
		}
		
		func (t *SimpleChaincode) viewPacketDtlsByPatientId(stub shim.ChaincodeStubInterface, args []string) pb.Response {
			logger.Info("inside viewPacketDtlsByPatientId---------------------",args)
			if len(args) != 1 {
				return shim.Error("Incorrect number of arguments. Expecting 1")
			}
			var packetIdBuffer bytes.Buffer
			var patientPacketArray [] string
			patientIdIterator, err := stub.GetQueryResult("{\"selector\":{\"docType\":\"PatientDetails\",\"patientId\":\"" + args[0] + "\"}}")
			logger.Info("patientDetail---------------------",patientIdIterator)
			if err != nil {
				return shim.Error(err.Error())
			}
			for patientIdIterator.HasNext() {
				resultQuery, resultError := patientIdIterator.Next()
				logger.Info("inside for loop---------------------",resultQuery)
				if resultError != nil {
					return shim.Error(resultError.Error())
				}
				var patientDetails PatientDetails
				unMarshalError := json.Unmarshal(resultQuery.Value, &patientDetails)
				if unMarshalError != nil {
					return shim.Error(unMarshalError.Error())
				}
				patientPacketArray=patientDetails.PacketInfo

				logger.Info("patientpacket array is --------------",patientPacketArray)
				for i:=0;i<len(patientPacketArray);i++ {
					packetIdIterator, err := stub.GetQueryResult("{\"selector\":{\"docType\":\"PacketDetails\",\"packetId\":\"" + patientPacketArray[i] + "\"}}")
					logger.Info("packet details for patinet is --------------",packetIdIterator)
					if err != nil {
						return shim.Error(err.Error())
					}
					defer packetIdIterator.Close()
					// buffer is a JSON array containing QueryResults
					
					packetIdBuffer.WriteString("[")
		
					packetIdAlreadyWritten := false
					for packetIdIterator.HasNext() {
						queryResponse, err := packetIdIterator.Next()
						logger.Info("inside create json------------",queryResponse)
						if err != nil {
							return shim.Error(err.Error())
						}
						// Add a comma before array members, suppress it for the first array member
						if packetIdAlreadyWritten == true {
							packetIdBuffer.WriteString(",")
						}
						packetIdBuffer.WriteString("{\"Key\":")
						packetIdBuffer.WriteString("\"")
						packetIdBuffer.WriteString(queryResponse.Key)
						packetIdBuffer.WriteString("\"")
		
						packetIdBuffer.WriteString(", \"Record\":")
						// Record is a JSON object, so we write as-is
						packetIdBuffer.WriteString(string(queryResponse.Value))
						packetIdBuffer.WriteString("}")
						packetIdAlreadyWritten = true
					}
					packetIdBuffer.WriteString("]")
				
				}
				
			}
			return shim.Success(packetIdBuffer.Bytes())
		}	
		
		func (t *SimpleChaincode) viewPacketDtlsByIssueDate(stub shim.ChaincodeStubInterface, args []string) pb.Response {
			if len(args) != 1 {
				return shim.Error("Incorrect number of arguments. Expecting 1")
			}
			packetIdIterator, err := stub.GetQueryResult("{\"selector\":{\"docType\":\"PacketDetails\",\"issueDate\":\"" + args[0] + "\"}}")
			if err != nil {
				return shim.Error(err.Error())
			}
		
			defer packetIdIterator.Close()
			// buffer is a JSON array containing QueryResults
			var packetIdBuffer bytes.Buffer
			packetIdBuffer.WriteString("[")
		
			packetIdAlreadyWritten := false
			for packetIdIterator.HasNext() {
				queryResponse, err := packetIdIterator.Next()
				if err != nil {
					return shim.Error(err.Error())
				}
				// Add a comma before array members, suppress it for the first array member
				if packetIdAlreadyWritten == true {
					packetIdBuffer.WriteString(",")
				}
				packetIdBuffer.WriteString("{\"Key\":")
				packetIdBuffer.WriteString("\"")
				packetIdBuffer.WriteString(queryResponse.Key)
				packetIdBuffer.WriteString("\"")
		
				packetIdBuffer.WriteString(", \"Record\":")
				// Record is a JSON object, so we write as-is
				packetIdBuffer.WriteString(string(queryResponse.Value))
				packetIdBuffer.WriteString("}")
				packetIdAlreadyWritten = true
			}
			packetIdBuffer.WriteString("]")
		
			return shim.Success(packetIdBuffer.Bytes())	
		}
		
		func (t *SimpleChaincode) trackPackageByPacketId(stub shim.ChaincodeStubInterface, args []string) pb.Response {
			if len(args) != 1 {
				return shim.Error("Incorrect number of arguments. Expecting 1")
			}
		
			var lotId string
			//to store the lotId for further querying harvest asset
			//var certId string
			
		
			var patientBuffer bytes.Buffer
			var purchaseOrderBuffer bytes.Buffer
			var lotBuffer bytes.Buffer
			var harvestBuffer bytes.Buffer
			var plantBuffer bytes.Buffer
		
		
			packetIdIterator, err := stub.GetQueryResult("{\"selector\":{\"docType\":\"PacketDetails\",\"packetId\":\"" + args[0] + "\"}}")
			if err != nil {
				return shim.Error(err.Error())
			}
		
			defer packetIdIterator.Close()
			// buffer is a JSON array containing QueryResults
			var packetIdBuffer bytes.Buffer
			packetIdBuffer.WriteString("[")
		
			packetIdAlreadyWritten := false
			for packetIdIterator.HasNext() {
				queryResponse, err := packetIdIterator.Next()
				if err != nil {
					return shim.Error(err.Error())
				}
				// Add a comma before array members, suppress it for the first array member
				if packetIdAlreadyWritten == true {
					packetIdBuffer.WriteString(",")
				}
				packetIdBuffer.WriteString("{\"Key\":")
				packetIdBuffer.WriteString("\"")
				packetIdBuffer.WriteString(queryResponse.Key)
				packetIdBuffer.WriteString("\"")
		
				packetIdBuffer.WriteString(", \"Record\":")
				// Record is a JSON object, so we write as-is
				packetIdBuffer.WriteString(string(queryResponse.Value))
				packetIdBuffer.WriteString("}")
				packetIdAlreadyWritten = true
			}
			packetIdBuffer.WriteString("]")
		
			patientIteratorPacId, errPacId := stub.GetQueryResult("{\"selector\":{\"docType\":\"PatientDetails\"}}")
			//for selecting a patient and getting packetInfo to match with the given packetId
			if errPacId != nil {
				return shim.Error(errPacId.Error())
			}
			var patientPacketArray [] string
			patientBuffer.WriteString("[")
			packetIdAlreadyWritten = false
			for patientIteratorPacId.HasNext() {
				resultQuery, resultError := patientIteratorPacId.Next()
				if resultError != nil {
					return shim.Error(resultError.Error())
				}
				var patientDetails PatientDetails
				unMarshalError := json.Unmarshal(resultQuery.Value, &patientDetails)
				if unMarshalError != nil {
					return shim.Error(unMarshalError.Error())
				}
				patientPacketArray=patientDetails.PacketInfo

				for i:=0;i<len(patientPacketArray);i++ {
		
					if args[0]==patientPacketArray[i] {
						if packetIdAlreadyWritten {
						patientBuffer.WriteString(",")
						}
						patientBuffer.WriteString("{\"Key\":")
						patientBuffer.WriteString("\"")
						patientBuffer.WriteString(resultQuery.Key)
						patientBuffer.WriteString("\"")
		
						patientBuffer.WriteString(", \"Record\":")
						// Record is a JSON object, so we write as-is
						patientBuffer.WriteString(string(resultQuery.Value))
						patientBuffer.WriteString("}")
						break
		
					}
				}
			}
			patientBuffer.WriteString("]")
			var purchaseOrderPacketIdArray[] string
			//to store the packet Array of a particular purchase order
			purchaseOrderIterator, errPurchaseOrder := stub.GetQueryResult("{\"selector\":{\"docType\":\"PurchaseOrderDetails\"}}")
			if errPurchaseOrder != nil {
				return shim.Error(errPurchaseOrder.Error())
			}
			packetIdAlreadyWritten = false
			purchaseOrderBuffer.WriteString("[")
			for purchaseOrderIterator.HasNext() {
				resultQuery, resultError := purchaseOrderIterator.Next()
				//from here
				if resultError != nil {
					return shim.Error(resultError.Error())
				}
				var purchaseOrderDetails PurchaseOrderDetails
				unMarshalError := json.Unmarshal(resultQuery.Value, &purchaseOrderDetails)
				if unMarshalError != nil {
					return shim.Error(unMarshalError.Error())
				}
				purchaseOrderPacketIdArray=purchaseOrderDetails.PacketInfo
		
				for i:=0;i<len(purchaseOrderPacketIdArray);i++ {
				//can just send the dispensary Id alone no need to send an entire asset
					if purchaseOrderPacketIdArray[i]==args[0] {
					//use break statement
					//no need to use like this although can be done save it for future reference	
						if packetIdAlreadyWritten {
						purchaseOrderBuffer.WriteString(",")
						}
				
						purchaseOrderBuffer.WriteString("{\"Key\":")
						purchaseOrderBuffer.WriteString("\"")
						purchaseOrderBuffer.WriteString(resultQuery.Key)
						purchaseOrderBuffer.WriteString("\"")
			
						purchaseOrderBuffer.WriteString(", \"Record\":")
						purchaseOrderBuffer.WriteString(string(resultQuery.Value))
						purchaseOrderBuffer.WriteString("}")
						packetIdAlreadyWritten = true
						break
						
						}
					}
				}
			purchaseOrderBuffer.WriteString("]")
			var lotPacketIdArray[] string
			lotIterator, errLot := stub.GetQueryResult("{\"selector\":{\"docType\":\"LotDetails\"}}")
			if errLot != nil {
				return shim.Error(errLot.Error())
			}
			packetIdAlreadyWritten = false
			lotBuffer.WriteString("[")
			for lotIterator.HasNext() {
				resultQuery, resultError := lotIterator.Next()
				//from here
				if resultError != nil {
					return shim.Error(resultError.Error())
				}
				var lotDetails LotDetails
				unMarshalError := json.Unmarshal(resultQuery.Value, &lotDetails)
				if unMarshalError != nil {
					return shim.Error(unMarshalError.Error())
				}
				lotPacketIdArray=lotDetails.PacketInfo
		
				for i:=0;i<len(lotPacketIdArray);i++ {
		
					if lotPacketIdArray[i]==args[0] {
						if packetIdAlreadyWritten {
						lotBuffer.WriteString(",")
						}		
						lotBuffer.WriteString("{\"Key\":")
						lotBuffer.WriteString("\"")
						lotBuffer.WriteString(resultQuery.Key)
						lotId=resultQuery.Key
						lotBuffer.WriteString("\"")
						lotBuffer.WriteString(", \"Record\":")
						lotBuffer.WriteString(string(resultQuery.Value))
						lotBuffer.WriteString("}")
						packetIdAlreadyWritten =true
						break
						
						}
					}
				}
			lotBuffer.WriteString("]")
			var harvestLotIdArray[] string
			//var certificateId string
			harvestIterator, errHarvest := stub.GetQueryResult("{\"selector\":{\"docType\":\"HarvestDetails\"}}")
			if errHarvest != nil {
				return shim.Error(errHarvest.Error())
			}
			packetIdAlreadyWritten = false
			harvestBuffer.WriteString("[")
			for harvestIterator.HasNext() {
				resultQuery, resultError := harvestIterator.Next()
				//from here
				if resultError != nil {
					return shim.Error(resultError.Error())
				}
				var harvestDetails HarvestDetails
				unMarshalError := json.Unmarshal(resultQuery.Value, &harvestDetails)
				if unMarshalError != nil {
					return shim.Error(unMarshalError.Error())
				}
				logger.Info("harvest details is",harvestDetails)
				harvestLotIdArray=harvestDetails.LotInfo
				//certificateId=harvestDetails.CertificateID
				for i:=0;i<len(harvestLotIdArray);i++ {
				logger.Info("inside harvest lot id is and harvest lot array --------------",lotId,harvestLotIdArray[i] )
					if harvestLotIdArray[i]==lotId {
						if packetIdAlreadyWritten == true {
						harvestBuffer.WriteString(",")
					 	}
				logger.Info("inside if condition --------------",resultQuery.Key,resultQuery.Value)

						harvestBuffer.WriteString("{\"Key\":")
						harvestBuffer.WriteString("\"")
						harvestBuffer.WriteString(resultQuery.Key)
						harvestBuffer.WriteString("\"")
			
						harvestBuffer.WriteString(", \"Record\":")
						harvestBuffer.WriteString(string(resultQuery.Value))
						harvestBuffer.WriteString("}")

						packetIdAlreadyWritten = true
						break
						
						}
					}
				}
			harvestBuffer.WriteString("]")
			// plantIterator, errPlant:= stub.GetQueryResult("{\"selector\":{\"docType\":\"PlantDetails\"}}")
			// if errPlant != nil {
			// 	return shim.Error(errPlant.Error())
			// }
			// packetIdAlreadyWritten = false
			plantBuffer.WriteString("[")
			// for plantIterator.HasNext() {
			// 	resultQuery, resultError := plantIterator.Next()
			// 	from here
			// 	if resultError != nil {
			// 		return shim.Error(resultError.Error())
			// 	}
			// 	var plantDetails PlantDetails
			// 	unMarshalError := json.Unmarshal(resultQuery.Value, &plantDetails)
			// 	if unMarshalError != nil {
			// 		return shim.Error(unMarshalError.Error())
			// 	}
			// 	certId=plantDetails.CertificateID
			// 		if certId==certificateId {

			// 			if packetIdAlreadyWritten {
			// 			plantBuffer.WriteString(",")
			// 			}
			// 			plantBuffer.WriteString("{\"Key\":")
			// 			plantBuffer.WriteString("\"")
			// 			plantBuffer.WriteString(resultQuery.Key)
			// 			plantBuffer.WriteString("\"")
			
			// 			plantBuffer.WriteString(", \"Record\":")
			// 			plantBuffer.WriteString(string(resultQuery.Value))
			// 			plantBuffer.WriteString("}")
			// 			packetIdAlreadyWritten = true

		
						
			// 			}
			// 		}
				plantBuffer.WriteString("]")
			var res = TrackResponse{PacketAsset: 		packetIdBuffer.String(),
						PatientAsset:		patientBuffer.String(),
						PurchaseOrderAsset: purchaseOrderBuffer.String(),
						LotAsset: 			lotBuffer.String(),
						HarvestAsset: 		harvestBuffer.String(),
						PlantAsset: 		plantBuffer.String()}
					finalResponse, marshalerr := json.Marshal(res)
					if marshalerr != nil {
						shim.Error(marshalerr.Error())
					}
			
			return shim.Success(finalResponse)
		
		
		
			}

		func (t *SimpleChaincode) addPatient(stub shim.ChaincodeStubInterface, args []string) pb.Response {
			logger.Info("inside addPatient---------------------",args)
			if len(args) != 12 {
					return shim.Error("Incorrect number of arguments. Expecting 12")
				}
			patientIterator, err := stub.GetQueryResult("{\"selector\":{\"docType\":\"Patient\",\"patientId\":\"" + args[0] + "\"}}")
			if err != nil {
				return shim.Error(err.Error())
			}
			logger.Info("iterated patient ID is",patientIterator)
			if patientIterator.HasNext() {
				logger.Info("inside patientiD iterator")
				var err error
				resultQuery, resultError := patientIterator.Next()
				if resultError != nil {
					return shim.Error(resultError.Error())
				}
				var patientDt Patient
				unMarshalError := json.Unmarshal(resultQuery.Value, &patientDt)
				if unMarshalError != nil {
					return shim.Error(unMarshalError.Error())
				}
				logger.Info("fetch patient Details---------------------",patientDt)
				var prescription=PatientPrescription{
					AssetType:		"patient_Prescription",
					Date:			args[6],
					Doses:			args[7],
					Type:			args[8],
					Grade:			args[9],
					Status:			"Pending",
					DoctorID:		args[10],
					Prescription:	args[11]}
					logger.Info("patient prescription is",prescription)
				// temp := `{"docType":"patient_Prescription",  "Date": "` + args[6] + `" , "Doses": "` + args[7] + `" , "Type": "` + args[8] + `" , "Grade": "`+args[9] +`" , "Status": "` + "Pending" + `", "DoctorID": "` + args[10] + `" , "prescription:"` + args[11] + ` "}`
				patientDt.Patient_Prescription=append(patientDt.Patient_Prescription,prescription)
				logger.Info("updated data is ---------------------",patientDt)
				var patient = Patient{
					AssetType:   	  	"Patient",
					PatientId:     		patientDt.PatientId,
					Name:   		  	patientDt.Name,
					SSN:  	  	  	  	patientDt.SSN,
					Age:	  	  	  	patientDt.Age,
					PhoneNo:		  	patientDt.PhoneNo,
					Dispensary:     	patientDt.Dispensary,
					Patient_Prescription: patientDt.Patient_Prescription}
					logger.Info("patient is---------------------",patient)
					patientDetailAsBytes, _ := json.Marshal(patient)
				err = stub.PutState(args[0], patientDetailAsBytes)
				logger.Info("error is",err)
				if err != nil {
					return shim.Error("Error while put state to update patient ")
				}
				patientIdIterator, err := stub.GetQueryResult("{\"selector\":{\"docType\":\"Patient\",\"patientId\":\"" + args[0] + "\"}}")
		if err != nil {
			return shim.Error(err.Error())
		}
		logger.Info("patient data is",patientIdIterator)
		for patientIdIterator.HasNext() {
			logger.Info("inside for loop  data  ----------------------")
			queryResponse, err := patientIdIterator.Next()
			if err != nil {
				return shim.Error(err.Error())
			}
			logger.Info("queryResponse is-------------------",queryResponse.Value)

		}
		logger.Info("patient Detailsa are ",patientIdIterator)
				return shim.Success([]byte("Prescription updated successfully for PatientID " + args[0] ))
			} else {
					logger.Info("inside else part")
					ssnIdIterator, err := stub.GetQueryResult("{\"selector\":{\"docType\":\"Patient\",\"SSN\":\"" + args[1] + "\"}}")
					if err != nil {
						return shim.Error(err.Error())
					}
					logger.Info("all ssn id's are",ssnIdIterator)
				if ssnIdIterator.HasNext() {
					return shim.Success([]byte("SSN exists"))
				} else {
					logger.Info("data with non existed patientId & SSN ")
					// var patient_pres= patient_Prescription {
					// Date: 			args[6],
					// Doses:			args[7],
					// Type:			args[8],
					// Grade:			args[9],
					// Status:			"Pending",
					// DoctorID:		args[10],
					// prescription:	args[11]}
				// var temp []patient_Prescription
				// temp=append(temp,patient_pres)
				// var temps []string
				var ptprescrisption []PatientPrescription
				var prescription=PatientPrescription{
					AssetType:		"patient_Prescription",
					Date:			args[6],
					Doses:			args[7],
					Type:			args[8],
					Grade:			args[9],
					Status:			"Pending",
					DoctorID:		args[10],
					Prescription:	args[11]}
					logger.Info("patient prescription is",prescription)
					ptprescrisption=append(ptprescrisption,prescription)
				// temp := `{"docType":"patient_Prescription",  "Date": "` + args[6] + `" , "Doses": "` + args[7] + `" , "Type": "` + args[8] + `" , "Grade": "`+args[9] +`" , "Status": "` + "Pending" + `", "DoctorID": "` + args[10] + `" , "prescription:"` + args[11] + ` "}`
				// temps=append(temps,temp)
				// logger.Info("temp data is",temp)
						var patientDt = Patient{
							AssetType:   "Patient",
							PatientId:     	args[0],
							Age:	  	  	  args[2],
							SSN:    	  	args[1],
							PhoneNo:		  args[3],
							Name:   		args[4],
							Dispensary:     args[5],
							Patient_Prescription:	ptprescrisption}
						logger.Info("data is ====================",patientDt)
						patientDetailAsBytes, _ := json.Marshal(patientDt)
						logger.Info("data as bytes",ptprescrisption)
						err = stub.PutState(args[0], patientDetailAsBytes)
						logger.Info("error is",err)
						if err != nil {
							return shim.Error("Error while put state of PatientDetail for adding new Patient.")
						}
						patientIdIterator, err := stub.GetQueryResult("{\"selector\":{\"docType\":\"Patient\",\"patientId\":\"" + args[0] + "\"}}")
		// patientIdIterator, err := stub.GetQueryResult("{\"selector\":{\"docType\":\"PatientDetail\"}}")
		if err != nil {
			return shim.Error(err.Error())
		}
		logger.Info("patient Detailsa are ",patientIdIterator)
		for patientIdIterator.HasNext() {
			logger.Info("inside for loop  data  ----------------------")
			queryResponse, err := patientIdIterator.Next()
			if err != nil {
				return shim.Error(err.Error())
			}
			logger.Info("queryResponse is-------------------",queryResponse.Value)

		}
			return shim.Success([]byte("Prescription added successfully for PatientID" + args[0] ))
				}
			}
			return shim.Success(nil)	
		}
	func (t *SimpleChaincode) getPatientDetails(stub shim.ChaincodeStubInterface) pb.Response {
		logger.Info("inside getPatientDetails---------------")
		patientIdIterator, err := stub.GetQueryResult("{\"selector\":{\"docType\":\"Patient\"}}")
		if err != nil {
			return shim.Error(err.Error())
		}
		defer patientIdIterator.Close()
		var patientBuffer bytes.Buffer
		patientBuffer.WriteString("[")
	
		patientAlreadyWritten := false
		for patientIdIterator.HasNext() {
			queryResponse, err := patientIdIterator.Next()
			if err != nil {
				return shim.Error(err.Error())
			}
			if patientAlreadyWritten == true {
				patientBuffer.WriteString(",")
			}
			patientBuffer.WriteString("{\"Key\":")
			patientBuffer.WriteString("\"")
			patientBuffer.WriteString(queryResponse.Key)
			patientBuffer.WriteString("\"")
	
			patientBuffer.WriteString(", \"Record\":")
			logger.Info("query response is",string(queryResponse.Value))
			patientBuffer.WriteString(string(queryResponse.Value))
			patientBuffer.WriteString("}")
			patientAlreadyWritten = true
		}
		patientBuffer.WriteString("]")
		return shim.Success(patientBuffer.Bytes())
	}
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		logger.Errorf("Error starting Simple chaincode: %s", err)
	}
}
