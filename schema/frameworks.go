package schema

// Framework represents a standard with a name, version, web link, and a list of controls
// All controls must have a name, framework slug, version, and at least one control
// This will be validated by the jsonschema
type Framework[T any] struct {
	// Name of the audit standard in long form (e.g. SOC 2, NIST CSF)
	Name string `json:"name" jsonschema:"minLength=1,description=name of the framework,example=SOC 2,example=NIST CSF"`
	// Framework is the shortname (slug) of the framework
	Framework string `json:"framework" jsonschema:"minLength=1,description=shortname (slug) of the framework,example=soc2,example=nist_csf"`
	// Version of the audit standard for the controls
	Version string `json:"version" jsonschema:"minLength=1,description=version of the framework,example=2017,example=1.1"`
	// WebLink to documentation for the standard
	WebLink string `json:"web_link,omitempty" jsonschema:"type=string,description=link to the documentation for the controls"`
	// Controls for the standard, minimum of 1 control is required for a valid standard
	Controls []Control[T] `json:"controls" jsonschema:"description=a list of controls for the framework"`
}

// Control is the fields that a control may have, all controls must have a ref code (the unique identifier) and a category
// all other fields are optional
// Framework specific metadata can be added to the metadata field using the appropriate type
type Control[T any] struct {
	// Name of the control
	Name string `json:"name,omitempty" jsonschema:"description=name of the control"`
	// Description of the control
	Description string `json:"description,omitempty" jsonschema:"description=short description of the control"`
	// RefCode is the unique identifier for the control
	RefCode string `json:"ref_code" jsonschema:"minLength=1,description=unique identifier for the control in the framework; sometimes referred to as a control number,example=CC1.1,example=AC-2"`
	// Category is the category of the control
	Category string `json:"category" jsonschema:"minLength=1,description=category of the control,example=Security,example=Availability"`
	// Subcategory is the subcategory of the control
	Subcategory string `json:"subcategory,omitempty" jsonschema:"description=subcategory of the control,example=System Operation,example=Control Environment"`
	// SubControls are the sub controls of the control
	SubControls []Control[T] `json:"sub_controls,omitempty" jsonschema:"description=sub-controls of the control if applicable"`
	// DTI is the difficulty to implement for the control, which can be set to easy, medium, or difficult
	DTI string `json:"dti,omitempty" jsonschema:"enum=easy,enum=medium,enum=difficult,description=difficulty to implement rating of the control"`
	// DTC is the difficulty to collect evidence for the control, which can be set to easy, medium, or difficult
	DTC string `json:"dtc,omitempty" jsonschema:"enum=easy,enum=medium,enum=difficult,description=difficulty to collect evidence rating of the control"`
	// Guidance is the guidance or suggested steps for implementing the control
	Guidance string `json:"guidance,omitempty" jsonschema:"description=guidance or suggested steps for implementing the control"`
	// Metadata is the metadata for the control
	Metadata T `json:"metadata,omitempty" jsonschema:"oneof_type=object;array,description=metadata for the control; unique to the framework"`
}
