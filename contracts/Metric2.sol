// SPDX-License-Identifier: MIT

pragma solidity ^0.8.7;

contract Metric2{
    // Metadata structure
	struct Metadata {
		string metricID; 
		string metricName; 
		string tenantID;
		uint noOfFormulas;
		string trustNetPK;
	}

    // Formula structure
	struct Formula {
		string formulaID; 
		address contractAddress;
		uint noOfValues;
		string activityID;
		string activityName; 
		string valueIDs;
	}

    // Value structure
	struct Value {
		string valueID;
		string valueName;
        string formulaID;
		string workflowID;
		string stageID;
		string stageName; 
		string keyName; 
		string tdpType;
		int bindingType;
		string artifactID;
		string primaryKeyRowID;
		string artifactTemplateName; 
		string fieldKey; 
		string fieldName; 
	}

    //Metadata setter
	Metadata metadata = Metadata("63367af765d44a7b142f7DevTest01", "Carbon Footprint", "77ce7ab0-c77e-11ec-b6ff-411ad42139cd", 2, "GBOK5LA4FABVOK76XOZGYHXDF5XYMS5FQJ7XZNP6TTAG6NYS766TO7EF");

    //Formula setters for each formula
    //Here the variable name should be formula name
    Formula Fertilizer_Emission = Formula("631a0b4ad9241a9374fConfig01", address(0xB861Ab9cA8caBe266078e9f25ba5ab99B3d2f67A), 2, "63367e1d218ee685c5e1a001", "53656564696e6720262047656d696e6174696f6e2066657274696c697a6572207573616765", "e1223109481cf739d19e6735d0236577, e1223109481cf739d19e6735d0236577");
    Formula Water_Emission = Formula("6360bc1c2e2e07a704babfa6", address(0xaA1Adc9cF23F234a3927Df60BD020a328Aeb9031), 1, "63367e1d218ee685c5e1a002", "53656564696e6720262047656d696e6174696f6e207761746572207573616765", "e1223109481cf739d19e6735d0236578");

    //Value setters for each value
    //Here the variable name should be value name
    //Values for formula 1
    Value FERTILIZER_EMISSION = Value("e1223109481cf739d19e6735d0236577", "FERTILIZER_EMISSION", "631a0b4ad9241a9374fConfig01" ,"61373288d9db6363906c7512", "100", "53656564696e6720262047656d696e6174696f6e2066657274696c697a657220757361676520e7a8aee381bee3818de383bbe799bae88abde882a5e69699e381aee4bdbfe38184e696b9", "e0b6b4e0b79ce0b784e0b79ce0b6bb2fe0ae89e0aeb0e0aeaee0af8d", "5", 2, "62299e34cc61b2d646056147", "testID", "46657274696c697a6572", "656d697373696f6e466163746f72", "46657274696c697a65722054797065");
    Value COMPOST_QUANTITY = Value("e1223109481cf739d19e6735d0236578", "FERTILIZER_EMISSION", "631a0b4ad9241a9374fConfig01", "61373288d9db6363906c7512", "100", "53656564696e6720262047656d696e6174696f6e2066657274696c697a657220757361676520e7a8aee381bee3818de383bbe799bae88abde882a5e69699e381aee4bdbfe38184e696b9", "e0b6b4e0b79ce0b784e0b79ce0b6bbe0b6b4e0b78ae2808de0b6bbe0b6b8e0b78fe0b6abe0b6ba2fe0ae89e0aeb0e0aeaee0af8de0ae85e0aeb3e0aeb5e0af81", "5", 1, "", "", "", "", "");

    //Values for formula 1
    Value WATER = Value("e1223109481cf739d19e6735d0236579", "Water", "6360bc1c2e2e07a704babfa6", "61373288d9db6363906c7512", "100", "53656564696e6720262047656d696e6174696f6e20776174657220757361676520e692ade7a8aee383bbe799bae88abde6b0b4e381aee4bdbfe794a8e9878f", "e0b6a2e0b6bde0b6b4e0b6bbe0b792e0b6b7e0b79de0b6a2e0b6b1e0b6ba2fe0aea4e0aea3e0af8de0aea3e0af80e0aeb0e0af8de0aea8e0af81e0ae95e0aeb0e0af8de0aeb5e0af81", "5", 1, "", "", "", "", "");

    //formula getters for each formula
    function getFertilizer_Emission() public view returns(string memory, address, uint, string memory, string memory, string memory){
        return (Fertilizer_Emission.formulaID, Fertilizer_Emission.contractAddress, Fertilizer_Emission.noOfValues, Fertilizer_Emission.activityID, Fertilizer_Emission.activityName, Fertilizer_Emission.valueIDs);
    }

    function getWater_Emission() public view returns(string memory, address, uint, string memory, string memory, string memory){
        return (Water_Emission.formulaID, Water_Emission.contractAddress, Water_Emission.noOfValues, Water_Emission.activityID, Water_Emission.activityName, Water_Emission.valueIDs);
    }

    //value getters for each value
    function getFERTILIZER_EMISSION() public view returns(Value memory){
        return FERTILIZER_EMISSION;
    }

    function getCOMPOST_QUANTITY() public view returns(Value memory){
        return COMPOST_QUANTITY;
    }

    function getWATER() public view returns(Value memory){
        return WATER;
    }
}