 /*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

/*
 * The sample smart contract for documentation topic:
 * Writing Your First Blockchain Application
 */

// Written by Xu Chen Hao
// Building on windows:
// 1. Install cygwin64 with gcc/g++
// 2. Set system Path env to C:\cygwin64\bin
// 3. Install golang v1.9.2 & set GOPATH env to D:\go
// 4. mkdir D:\go\src\sacc\ & put this chaincoe in there
// 5. go get -u --tags nopkcs11 github.com/hyperledger/fabric/core/chaincode/shim
// 6. go build --tags nopkcs11
package main

/* Imports
 * 4 utility libraries for formatting, handling bytes, reading and writing JSON, and string manipulation
 * 2 specific Hyperledger Fabric specific libraries for Smart Contracts
 */
import (
	"encoding/json"
	"encoding/hex"
	"fmt"
	"strings"
	"math/rand"
	"strconv"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

// Define the car structure, with 4 properties.  Structure tags are used by encoding/json library
//human ID
type Human struct {
	
	ID            string `json:"身份证号"`
	Sex           string `json:"性别"`
	Name          string `json:"姓名"`
	Date          string `json:"出生日期"`
	FatherName    string `json:"父亲姓名"`
	FatherID      string `json:"父亲身份证号"`
	MotherName    string `json:"母亲姓名"`
	MotherID      string `json:"母亲身份证号"`
	MarryState    string `json:"婚姻状态"`
	SpouseName    string `json:"配偶姓名"`
	SpouseID      string `json:"配偶身份证号"`
	Marry_Cert    string `json:"结婚证书"`
	ChildID  [10] string `json:"子女身份证号"`
	ChildName[10] string `json:"子女姓名"`
	NewChild [10] string `json:"子女出生证明"`
}

//birth cert
type Birth struct {

	BirthID      string `json:"出生证书编号"`
	Date         string `json:"出生日期"`
	Sex          string `json:"性别"`
	FatherName   string `json:"父亲姓名"`
	FatherID     string `json:"父亲身份证号"`
	MotherName   string `json:"母亲姓名"`
	MotherID     string `json:"母亲身份证号"`
	HosptialID   string `json:"接生机构"`
}

//marry card
type Marry_Card struct {

	Marry_Cert     string `json:"证书编号"`
	State          string `json:"状态"`
	Husband_Name   string `json:"丈夫姓名"`
	Husband_ID     string `json:"丈夫身份证号"`
	Wife_Name      string `json:"妻子姓名"`
	Wife_ID        string `json:"妻子身份证号"`
	Date           string `json:"登记日期"`
}

//marry check 
type Marry_Check struct{
	CheckID  		 string `json:"审查编号"`
	Husband_Name     string `json:"丈夫姓名"`
	Husband_ID       string `json:"丈夫身份证号"`
	HusbandState     string `json:"丈夫婚姻状态"`
	Wife_Name        string `json:"妻子姓名"`
	Wife_ID          string `json:"妻子身份证号"`
	WifeState        string `json:"妻子婚姻状态"`
	Check [6]        string `json:"判断结果"`
}

type Creat_Check struct{
	CheckID            string `json:"审查编号"`
	FatherName         string `json:"父亲姓名"`
	FatherID           string `json:"父亲身份证号"`
	MotherName         string `json:"母亲姓名"`
	MotherID           string `json:"母亲身份证号"`
	Marry_Cert         string `json:"父母结婚证书编号"`
	BirthID            string `json:"出生证书编号"`
	BirthDate          string `json:"出生日期"`
	Sex                string `json:"性别"`
	HosptialID         string `json:"接生机构"`
	Check [9]          string `json:"判断结果"`
}

type Divorce_Check struct{
	CheckID  		 string `json:"审查编号"`
	Husband_Name     string `json:"丈夫姓名"`
	Husband_ID       string `json:"丈夫身份证号"`
	Wife_Name        string `json:"妻子姓名"`
	Wife_ID          string `json:"妻子身份证号"`
	Marry_Cert       string `json:"结婚证书编号"`
	Check [5]        string `json:"判断结果"`
}
/*
 * The Init method is called when the Smart Contract "fabcar" is instantiated by the blockchain network
 * Best practice is to have any Ledger initialization in separate function -- see initLedger()
 */
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 * The Invoke method is called as a result of an application request to run the Smart Contract "fabcar"
 * The calling application program has also specified the particular smart contract function to be called, with arguments
 */
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "queryID" {
		return s.queryID(APIstub,args)
	}else if function == "createBirth" {
		return s.createBirth(APIstub, args)
	}else if function == "createHuman" {
		return s.createHuman(APIstub, args)
	}else if function == "marry" {
		return s.marry(APIstub, args)
	}else if function == "marryCheck" {
		return s.marryCheck(APIstub, args)
	}else if function == "divorceCheck" {
		return s.divorceCheck(APIstub, args)
	}else if function == "createCheck" {
		return s.createCheck(APIstub, args)
	}else if function == "initLedger" {
		return s.initLedger(APIstub)
	}else if function == "addInter" {
		return s.addInter(APIstub,args)
	}
	return shim.Error("Invalid Smart Contract function name.")
	
}	

func (s *SmartContract) queryID(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	humanAsBytes, err := APIstub.GetState(args[0])
	var human Human;
	err = json.Unmarshal(humanAsBytes,&human)//反序列化
	if err != nil {
		return shim.Error("{\"Error\":\"Failed to decode JSON of: " + string(humanAsBytes)+ "\" to Human}")
	}
   return shim.Success(humanAsBytes)
	
}

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	var humanA Human
	humanA.ID       = "110105199409026676"
	humanA.Sex      = "男"
	humanA.Name     = "李雷雷"
	humanA.FatherID = "110105197003025376"
	humanA.MotherID = "110105197302055386"
	humanA.ChildID[0] = "0"
 	humanA.NewChild[0] = "0"


	var humanB Human
	humanB.ID       = "110105199409026686"
	humanB.Sex      = "女"
	humanB.Name     = "韩梅梅"
	humanB.FatherID = "110105197107025376"
	humanB.MotherID = "110105197303055386"
	humanB.ChildID[0] = "0"
	humanB.NewChild[0] = "0"

	var humanC Human
	humanC.ID       = "110105199409026656"
	humanC.Sex      = "男"
	humanC.Name     = "王雷雷"
	humanC.FatherID = "110105197003025376"
	humanC.MotherID = "110105197302055386"
	humanC.ChildID[0] = "0"
	humanC.NewChild[0] = "0"


	var humanD Human
	humanD.ID       = "110105199409026646"
	humanD.Sex      = "女"
	humanD.Name     = "张梅梅"
	humanD.FatherID = "110105197107025376"
	humanD.MotherID = "110105197303055386"
	humanD.ChildID[0] = "0"
	humanD.NewChild[0] = "0"

	
	humanAsBytes, _ := json.Marshal(humanA)
	APIstub.PutState(humanA.ID, humanAsBytes)

	humanBAsBytes, _ := json.Marshal(humanB)
	APIstub.PutState(humanB.ID, humanBAsBytes)

	humanCsBytes, _ := json.Marshal(humanC)
	APIstub.PutState(humanC.ID, humanCsBytes)

	humanDAsBytes, _ := json.Marshal(humanD)
	APIstub.PutState(humanD.ID, humanDAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) createBirth(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	//5 paramtes father,mother,childsex,birhdate,hospitalID
	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}
	//whether father is sxisted
	FatherAsBytes, err := APIstub.GetState(args[0])
	var father Human;
	err = json.Unmarshal(FatherAsBytes,&father)//反序列化
	if err != nil {
		return shim.Error("{\"Error\":\"Failed to decode JSON of: " + string(FatherAsBytes)+ "\" to Human}")
	}
	//whether mother is sxisted
	MotherAsBytes, err := APIstub.GetState(args[1])
	var mother Human;
	err = json.Unmarshal(MotherAsBytes,&mother)//反序列化
	if err != nil {
		return shim.Error("{\"Error\":\"Failed to decode JSON of: " + string(MotherAsBytes)+ "\" to Human}")
	}

	//whether they are couples
	if 0  != (strings.Compare(father.SpouseID,mother.ID)){
		return shim.Error("{\"Error\":\"They are not couples }")
	}
	if 0  != (strings.Compare(mother.SpouseID,father.ID)){
		return shim.Error("{\"Error\":\"They are not couples }")
	}

	fnum,err := strconv.Atoi(father.NewChild[0])
	if fnum > 1{
		return shim.Error("{\"Error\":\"They are have enough children}")
	}
	mnum,err := strconv.Atoi(mother.ChildID[0])
	if mnum > 1{
		return shim.Error("{\"Error\":\"They are have enough children}")
	}

	//create birth certs
	var birth Birth;
	//timestamp := time.Now().Unix()
	//tm := time.Unix(timestamp, 0)
	rd := strconv.Itoa(rand.Intn(100))
	str := strings.Join([]string{args[3],rd},"")
	hashstr := hex.EncodeToString([]byte(str))

	birth.BirthID  = hashstr[0:18]
	birth.Sex      = args[2]
	birth.Date     = args[3]
	birth.FatherID = father.ID
	birth.FatherName = father.Name
	birth.MotherID = mother.ID
	birth.MotherName = mother.Name
	birth.HosptialID = args[4]
	birthAsBytes, _ := json.Marshal(birth)
	APIstub.PutState(birth.BirthID, birthAsBytes)

	//connected to the parents
	father.NewChild[0] = strconv.Itoa(fnum+1)
	father.NewChild[fnum+1] = birth.BirthID
	fatherAsBytes, _ := json.Marshal(father)
	APIstub.PutState(father.ID, fatherAsBytes)

	mother.NewChild[0] = strconv.Itoa(mnum+1)
	mother.NewChild[mnum+1] = birth.BirthID
	motherAsBytes, _ := json.Marshal(mother)
	APIstub.PutState(mother.ID, motherAsBytes)
	return shim.Success(birthAsBytes)
}

func (s *SmartContract) createCheck(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	//3 paramtes father or motherID ,1flow father 2 flow mother,name
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}
	var check Creat_Check

	//whether father is sxisted
	FatherAsBytes, err := APIstub.GetState(args[0])
	var father Human;
	err = json.Unmarshal(FatherAsBytes,&father)//反序列化
	if err != nil {
		check.Check[0] = "0"
		check.Check[1] = "0"
		check.FatherName   = "无"      
		check.FatherID     = "无"
	}else{
		check.Check[0] = "1"
		check.Check[1] = "1"
		check.FatherName   = father.Name      
		check.FatherID     = father.ID 
	}
	//whether mother is sxisted
	MotherAsBytes, err := APIstub.GetState(father.SpouseID)
	var mother Human;
	err = json.Unmarshal(MotherAsBytes,&mother)//反序列化
	if err != nil {
		check.Check[2] = "0"
		check.Check[3] = "0"
		check.MotherName   = "无"      
		check.MotherID     = "无" 
	}else{
		check.Check[2] = "1"
		check.Check[3] = "1"
		check.MotherName   = mother.Name      
		check.MotherID     = mother.ID 
	}

	//whether they are couples
	//whether married
	if 0 == len(father.SpouseID) {
		check.Check[4] = "0"
		check.Marry_Cert = "不是夫妻"
	}else if 0 == len(mother.SpouseID) {
		check.Check[4] = "0"
		check.Marry_Cert = "不是夫妻"
	}else if 0  != (strings.Compare(father.Marry_Cert,mother.Marry_Cert)){
		check.Check[4] = "0"
		check.Marry_Cert = "不是夫妻"
	}else {
	check.Check[4] = "1"
	check.Marry_Cert = father.Marry_Cert
	}
	// //get the child birth cert
	 cnum,err := strconv.Atoi(father.NewChild[0])
	// if fnum > cnum{
	// 	check.Check[5] = "0"
	// 	check.Check[6] = "0"
	// 	check.Check[7] = "0"
	// 	check.Check[8] = "0"
	// }else{

	ChildAsBytes, err := APIstub.GetState(father.NewChild[cnum])
	var child Birth
	err = json.Unmarshal(ChildAsBytes,&child)//反序列化
	if err != nil {
		check.Check[5] = "0"
		check.Check[6] = "0"
		check.Check[7] = "0"
		check.Check[8] = "0"
		check.BirthID      = "无"    
		check.BirthDate    = "无"      
		check.Sex          = "无"      
		check.HosptialID   = "无"
	}else{
		check.Check[5] = "1"
		check.Check[6] = "1"
		check.Check[7] = "1"
		check.Check[8] = "1"
		check.BirthID      = child.BirthID     
		check.BirthDate    = child.Date      
		check.Sex          = child.Sex      
		check.HosptialID   = child.HosptialID
	}

		str := strings.Join([]string{father.ID,mother.ID},"")
		hashstr := hex.EncodeToString([]byte(str))
		check.CheckID  =  hashstr[0:18]
		checkAsBytes, _ := json.Marshal(check)
	    APIstub.PutState(check.CheckID, checkAsBytes)    

		return shim.Success(checkAsBytes)
	}


func (s *SmartContract) createHuman(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	//3 paramtes father or motherID ,1flow father 2 flow mother,name
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}
	//whether father is sxisted
	FatherAsBytes, err := APIstub.GetState(args[0])
	var father Human;
	err = json.Unmarshal(FatherAsBytes,&father)//反序列化
	if err != nil {
		return shim.Error("{\"Error\":\"Failed to decode JSON of: " + string(FatherAsBytes)+ "\" to Human}")
	}
	//whether mother is sxisted
	MotherAsBytes, err := APIstub.GetState(father.SpouseID)
	var mother Human;
	err = json.Unmarshal(MotherAsBytes,&mother)//反序列化
	if err != nil {
		return shim.Error("{\"Error\":\"Failed to decode JSON of: " + string(MotherAsBytes)+ "\" to Human}")
	}

	//whether they are couples
	if 0  != (strings.Compare(father.SpouseID,mother.ID)){
		return shim.Error("{\"Error\":\"They are not couples }")
	}
	if 0  != (strings.Compare(mother.SpouseID,father.ID)){
		return shim.Error("{\"Error\":\"They are not couples }")
	}

	//whether more children
	fnum,err := strconv.Atoi(father.ChildID[0])
	if fnum > 2{
		return shim.Error("{\"Error\":\"They are have enough children}")
	}
	mnum,err := strconv.Atoi(mother.ChildID[0])
	if mnum > 2{
		return shim.Error("{\"Error\":\"They are have enough children}")
	}
	
	//get the child birth cert
	cnum,err := strconv.Atoi(father.NewChild[0])
	if fnum > cnum{
		return shim.Error("{\"Error\":\"The birth cert over time}")
	}

	ChildAsBytes, err := APIstub.GetState(father.NewChild[cnum])
	var child Birth
	err = json.Unmarshal(ChildAsBytes,&child)//反序列化
	if err != nil {
		return shim.Error("{\"Error\":\"Failed to decode JSON of birth cert}")
	}

	//create new human
	var newhuman Human
	newhuman.Sex      = child.Sex
	newhuman.Name     = args[2]
	newhuman.Date     = child.Date
	newhuman.MarryState = "未婚"
	newhuman.FatherID = father.ID
	newhuman.FatherName = father.Name
	newhuman.MotherID = mother.ID
	newhuman.MotherName = mother.Name
	newhuman.ChildID[0] = "0"
	newhuman.NewChild[0] = "0"
	
	if 0 == (strings.Compare("1",args[1])){
		address := father.ID[0:6]
		date := child.Date
		if 0 == (strings.Compare("男",child.Sex)){
			newhuman.ID = strings.Join([]string{address,date,"123",strconv.Itoa(rand.Intn(9))},"")
		}
		if 0 != (strings.Compare("男",child.Sex)){
			newhuman.ID = strings.Join([]string{address,date,"122",strconv.Itoa(rand.Intn(9))},"")
		}
	}

	if 0 == (strings.Compare("2",args[1])){
		address := mother.ID[0:6]
		date := child.Date
		if 0 == (strings.Compare("男",child.Sex)){
			newhuman.ID = strings.Join([]string{address,date,"123",strconv.Itoa(rand.Intn(9))},"")
		}
		if 0 != (strings.Compare("男",child.Sex)){
			newhuman.ID = strings.Join([]string{address,date,"122",strconv.Itoa(rand.Intn(9))},"")
		}
	}
	newhumanAsBytes, _ := json.Marshal(newhuman)
	APIstub.PutState(newhuman.ID, newhumanAsBytes)


	//become father
	fchild := strconv.Itoa(fnum+1)
	father.ChildID[0] = fchild
	father.ChildID[fnum+1] = newhuman.ID
	father.ChildName[fnum+1] = newhuman.Name
	fatherAsBytes, _ := json.Marshal(father)
	APIstub.PutState(father.ID, fatherAsBytes)
	//become mother
	mchild := strconv.Itoa(mnum+1)
	mother.ChildID[0] = mchild
	mother.ChildID[mnum+1] = newhuman.ID
	mother.ChildName[mnum+1] = newhuman.Name
	motherAsBytes, _ := json.Marshal(mother)
	APIstub.PutState(mother.ID, motherAsBytes)

	return shim.Success(newhumanAsBytes)
}

func (s *SmartContract) marryCheck(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	//2paramtes husbandID ,wifeID
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	var check Marry_Check
	//whether husband is exitd
	husbandAsBytes, err := APIstub.GetState(args[0])
	var husband Human;
	err = json.Unmarshal(husbandAsBytes,&husband)//反序列化
	if err != nil {
		check.Check[0] = "0"
		check.Check[1] = "0"
		check.Check[2] = "0"
		check.Husband_Name = "无"    
	    check.Husband_ID   = "无"
	}else{
		check.Check[0] = "1"
		check.Check[1] = "1"
		check.Husband_Name = husband.Name    
		check.Husband_ID   = husband.ID 
	}
	//whether wife is exitd
	wifeAsBytes, err := APIstub.GetState(args[1])
	var wife Human;
	err = json.Unmarshal(wifeAsBytes,&wife)//反序列化
	if err != nil {
		check.Check[3] = "0"
		check.Check[4] = "0"
		check.Check[5] = "0"
		check.Wife_Name    = "无"   
		check.Wife_ID      = "无"   
	}else{
		check.Check[3] = "1"
		check.Check[4] = "1"
		check.Wife_Name    = wife.Name    
		check.Wife_ID      = wife.ID  
	}
	//whether married
	if 0 != len(husband.SpouseID) {
		check.Check[2] = "0"
		check.HusbandState = "已婚"
	}else{
		check.Check[2] = "1"
		check.HusbandState = "未婚"
	}
	if 0 != len(wife.SpouseID) {
		check.Check[5] = "0"
		check.WifeState = "已婚"
	}else{
		check.Check[5] = "1"
		check.WifeState = "未婚"
	}

	str := strings.Join([]string{husband.ID,wife.ID},"")
	hashstr := hex.EncodeToString([]byte(str))

	check.CheckID      = hashstr[0:18]
	      
	  
  
	

	checkAsBytes, _ := json.Marshal(check)
	APIstub.PutState(check.CheckID, checkAsBytes)

	return shim.Success(checkAsBytes)
}

func (s *SmartContract) divorceCheck(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	//3paramtes husbandID ,wifeID
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	var check Divorce_Check
	//whether husband is exitd
	husbandAsBytes, err := APIstub.GetState(args[0])
	var husband Human;
	err = json.Unmarshal(husbandAsBytes,&husband)//反序列化
	if err != nil {
		check.Check[0] = "0"
		check.Check[1] = "0"
	}
	check.Check[0] = "1"
	check.Check[1] = "1"
	//whether wife is exitd
	wifeAsBytes, err := APIstub.GetState(args[1])
	var wife Human;
	err = json.Unmarshal(wifeAsBytes,&wife)//反序列化
	if err != nil {
		check.Check[2] = "0"
		check.Check[3] = "0"
	}
	check.Check[2] = "1"
	check.Check[3] = "1"
	//whether married
	if 0 == len(husband.SpouseID) {
		check.Check[4] = "0"
		check.Marry_Cert = "不是夫妻"
	}else if 0 == len(wife.SpouseID) {
		check.Check[4] = "0"
		check.Marry_Cert = "不是夫妻"
	}else if 0  != (strings.Compare(husband.SpouseID,wife.SpouseID)){
		check.Check[4] = "0"
		check.Marry_Cert = "不是夫妻"
	}else {
	check.Check[4] = "1"
	check.Marry_Cert = husband.Marry_Cert
	}

	str := strings.Join([]string{husband.ID,wife.ID},"")
	hashstr := hex.EncodeToString([]byte(str))

	check.CheckID      = hashstr[0:18]
	check.Husband_Name = husband.Name    
	check.Husband_ID   = husband.ID       
	check.Wife_Name    = wife.Name    
	check.Wife_ID      = wife.ID    
	 
	checkAsBytes, _ := json.Marshal(check)
	APIstub.PutState(check.CheckID, checkAsBytes)

	return shim.Success(checkAsBytes)
}

func (s *SmartContract) marry(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	//3paramtes husbandID ,wifeID,date
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}
	//whether husband is exitd
	husbandAsBytes, err := APIstub.GetState(args[0])
	var husband Human;
	err = json.Unmarshal(husbandAsBytes,&husband)//反序列化
	if err != nil {
		return shim.Error("{\"Error\":\"Failed to decode JSON of husband}")
	}
	//whether wife is exitd
	wifeAsBytes, err := APIstub.GetState(args[1])
	var wife Human;
	err = json.Unmarshal(wifeAsBytes,&wife)//反序列化
	if err != nil {
		return shim.Error("{\"Error\":\"Failed to decode JSON of wife}")
	}
	//whether married
	if 0 != len(husband.SpouseID) {
		return shim.Error("{\"Error\":\"Failed to married")
	}
	if 0 != len(wife.SpouseID) {
		return shim.Error("{\"Error\":\"Failed to married")
	}

	//generate marry id
	rd := strconv.Itoa(rand.Intn(100))
	str := strings.Join([]string{args[2],rd},"")
	hashstr := hex.EncodeToString([]byte(str))
	marry_cert_id  := strings.Join([]string{"J110101",args[2][0:4],hashstr[0:6]},"-")
	
	//become husband
	husband.SpouseID = wife.ID
	husband.SpouseName = wife.Name
	husband.MarryState = "已婚"
	husband.Marry_Cert = marry_cert_id
	husbandAsBytes, _ = json.Marshal(husband)
	APIstub.PutState(args[0], husbandAsBytes)
	//become wife
	wife.SpouseID = husband.ID
	wife.SpouseName = husband.Name
	wife.MarryState = "已婚"
	wife.Marry_Cert = marry_cert_id
	wifeAsBytes, _ = json.Marshal(wife)
	APIstub.PutState(args[1], wifeAsBytes)

	var card Marry_Card
	card.Marry_Cert = marry_cert_id
	card.State = "结婚"
	card.Husband_ID = husband.ID
	card.Husband_Name = husband.Name
	card.Wife_ID = wife.ID
	card.Wife_Name = wife.Name
	card.Date = strings.Join([]string{args[2][0:4],args[2][4:6],args[2][5:7]},"-")
	MarryAsBytes, _ := json.Marshal(card)
	APIstub.PutState(card.Marry_Cert, MarryAsBytes)

	return shim.Success(MarryAsBytes)
}

func (s *SmartContract) divorce(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}
	//whether husband is exitd
	husbandAsBytes, err := APIstub.GetState(args[0])
	var husband Human;
	err = json.Unmarshal(husbandAsBytes,&husband)//反序列化
	if err != nil {
		return shim.Error("{\"Error\":\"Failed to decode JSON of husband}")
	}
	//whether wife is exitd
	wifeAsBytes, err := APIstub.GetState(args[1])
	var wife Human;
	err = json.Unmarshal(wifeAsBytes,&wife)//反序列化
	if err != nil {
		return shim.Error("{\"Error\":\"Failed to decode JSON of wife}")
	}
	//whether they are couples
	if 0  != (strings.Compare(husband.SpouseID,wife.ID)){
		return shim.Error("{\"Error\":\"They are not couples ")
	}
	if 0  != (strings.Compare(wife.SpouseID,husband.ID)){
		return shim.Error("{\"Error\":\"They are not couples ")
	}

	//change Marry card
	var card Marry_Card
	cardAsBytes, err := APIstub.GetState(husband.Marry_Cert)
	err = json.Unmarshal(cardAsBytes,&card)//反序列化
	if err != nil {
		return shim.Error("{\"Error\":\"Failed to decode JSON of Marry_Card}")
	}
	card.State = "离婚"
	cardAsBytes, _ = json.Marshal(card)
	APIstub.PutState(husband.Marry_Cert, cardAsBytes)


	//change husband spouse
	husband.SpouseID = ""
	husband.Marry_Cert = ""
	husbandAsBytes, _ = json.Marshal(husband)
	APIstub.PutState(args[0], husbandAsBytes)
	//change wife spouse
	wife.SpouseID = ""
	wife.Marry_Cert = ""
	wifeAsBytes, _ = json.Marshal(wife)
	APIstub.PutState(args[1], wifeAsBytes)


	return shim.Success(nil)
}

func (s *SmartContract) addInter(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}
	var humanA Human
	humanA.ID       = args[0]
	humanA.Sex      = args[1]
	humanA.Name     = args[2]
	humanA.ChildID[0] = "0"
 	humanA.NewChild[0] = "0"

	humanAsBytes, _ := json.Marshal(humanA)
	APIstub.PutState(humanA.ID, humanAsBytes)

	return shim.Success(nil)
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}