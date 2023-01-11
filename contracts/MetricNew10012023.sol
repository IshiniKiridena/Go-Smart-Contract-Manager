// SPDX-License-Identifier: MIT

pragma solidity ^0.8.7;

contract Metric1_63367af765d44a7b142f7DevTest01a { 

	// Metadata structure
	struct Metadata {
		string metricID; 
		string metricName; // converted value to bytes
		string tenantID;
		uint noOfFormulas;
		string trustNetPK;
	}

	// Formula structure
	struct Formula {
		string formulaID; // actual formula ID
		address contractAddress;
		uint noOfValues;
		string activityID;
		string activityName; // converted value to bytes
		string valueIDs;
	}

	// Value structure
	struct Value {
		string valueID;
		string valueName;
		string workflowID;
		string stageID;
		string stageName; // converted value to bytes
		string keyName; // converted value to bytes
		string tdpType;
		int bindingType;
		string artifactID;
		string primaryKeyRowID;
		string artifactTemplateName; // converted value to bytes
		string fieldKey; // converted value to bytes
		string fieldName; // converted value to bytes
	}

    // Pivot Field structure 
    struct PivotField {
        string name;
        string key; // converted value to bytes
        string field; 
        string condition;
        string value;
        string artifactTemplateId;
        string artifactDataId;
        string formulaId;     
    }

	// Array to store all the values
	Value[] private allValues;

	// Array to store all the formulas
	Formula[] private allFormulas;

    // Array to store the pivot fields
    PivotField[] private pivotFields;

	// Metadata declaration
	Metadata metadata = Metadata("63367af765d44a7b142f7DevTest01a", "Carbon Footprint", "77ce7ab0-c77e-11ec-b6ff-411ad42139cd", 2, "GBOK5LA4FABVOK76XOZGYHXDF5XYMS5FQJ7XZNP6TTAG6NYS766TO7EF");

	// function to add details
	function addDetails() public {
		// add formula 1
		allFormulas.push(Formula("631a0b4ad9241a9374fConfig01", address(0xB861Ab9cA8caBe266078e9f25ba5ab99B3d2f67A), 2, "63367e1d218ee685c5e1a001", "53656564696e6720262047656d696e6174696f6e2066657274696c697a6572207573616765", "e1223109481cf739d19e6735d0236577, e1223109481cf739d19e6735d0236577"));
		// add value 1 for formula 1
		allValues.push(Value("e1223109481cf739d19e6735d0236577", "FERTILIZER_EMISSION", "61373288d9db6363906c7512", "100", "53656564696e6720262047656d696e6174696f6e2066657274696c697a657220757361676520e7a8aee381bee3818de383bbe799bae88abde882a5e69699e381aee4bdbfe38184e696b9", "e0b6b4e0b79ce0b784e0b79ce0b6bb2fe0ae89e0aeb0e0aeaee0af8d", "5", 2, "62299e34cc61b2d646056147", "testID", "46657274696c697a6572", "656d697373696f6e466163746f72", "46657274696c697a65722054797065"));
        // add value 2 for formula 1
		allValues.push(Value("e1223109481cf739d19e6735d0236578", "FERTILIZER_EMISSION", "61373288d9db6363906c7512", "100", "53656564696e6720262047656d696e6174696f6e2066657274696c697a657220757361676520e7a8aee381bee3818de383bbe799bae88abde882a5e69699e381aee4bdbfe38184e696b9", "e0b6b4e0b79ce0b784e0b79ce0b6bbe0b6b4e0b78ae2808de0b6bbe0b6b8e0b78fe0b6abe0b6ba2fe0ae89e0aeb0e0aeaee0af8de0ae85e0aeb3e0aeb5e0af81", "5", 1, "", "", "", "", ""));
        // add pivot fields formula 1
        pivotFields.push(PivotField("Selyn", "e0b6afe0b78ae0b780e0b792e0b6ade0b793e0b6bae0b792e0b69ae0b783e0b790e0b6b4e0b6bae0b794e0b6b8e0b78ae0b69ae0b6bbe0b7942fe0ae87e0aeb0e0aea3e0af8de0ae9fe0aebee0aeaee0af8de0aea8e0aebfe0aeb2e0af88e0ae9ae0aeaae0af8de0aeb3e0af88e0aeafe0aeb0e0af8d", "supplierName", "EQUAL", "638880305f7ec6ea4cd8398e", "62299e34cc61b2d646056147", "Selyn Exporters (Pvt) Limited", "631a0b4ad9241a9374fConfig01"));

		// add formula 2
		allFormulas.push(Formula("6360bc1c2e2e07a704babfa6", address(0xaA1Adc9cF23F234a3927Df60BD020a328Aeb9031), 1, "63367e1d218ee685c5e1a002", "53656564696e6720262047656d696e6174696f6e207761746572207573616765", "e1223109481cf739d19e6735d0236578"));
		// add value 1 for formula 2
		allValues.push(Value("e1223109481cf739d19e6735d0236579", "Water", "61373288d9db6363906c7512", "100", "53656564696e6720262047656d696e6174696f6e20776174657220757361676520e692ade7a8aee383bbe799bae88abde6b0b4e381aee4bdbfe794a8e9878f", "e0b6a2e0b6bde0b6b4e0b6bbe0b792e0b6b7e0b79de0b6a2e0b6b1e0b6ba2fe0aea4e0aea3e0af8de0aea3e0af80e0aeb0e0af8de0aea8e0af81e0ae95e0aeb0e0af8de0aeb5e0af81", "5", 1, "", "", "", "", ""));

    }

	// Getter for Formulas
	function getFormulaDetails() public view returns (Formula[] memory) {
		Formula[] memory formulaList = allFormulas;
		return formulaList;
	}

    // Getter for Values
    function getValueDetails() public view returns (Value[] memory) {
        Value[] memory valueList = allValues;
        return valueList;
    }

    // Getter for Pivot Fields
    function getPivotFields() public view returns (PivotField[] memory) {
        PivotField[] memory pivotFieldsList = pivotFields;
        return pivotFieldsList;
    }
}