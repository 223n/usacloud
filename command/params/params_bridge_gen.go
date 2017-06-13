// Code generated by 'github.com/sacloud/usacloud/tools/gen-input-models'; DO NOT EDIT

package params

import (
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

// ListBridgeParam is input parameters for the sacloud API
type ListBridgeParam struct {
	Name       []string
	Id         []int64
	From       int
	Max        int
	Sort       []string
	OutputType string
	Column     []string
	Quiet      bool
	Format     string
	FormatFile string
}

// NewListBridgeParam return new ListBridgeParam
func NewListBridgeParam() *ListBridgeParam {
	return &ListBridgeParam{}
}

// Validate checks current values in model
func (p *ListBridgeParam) Validate() []error {
	errors := []error{}
	{
		errs := validateConflicts("--name", p.Name, map[string]interface{}{

			"--id": p.Id,
		})
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Bridge"].Commands["list"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateConflicts("--id", p.Id, map[string]interface{}{

			"--name": p.Name,
		})
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := schema.ValidateInStrValues("json", "csv", "tsv")
		errs := validator("--output-type", p.OutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateOutputOption(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *ListBridgeParam) GetResourceDef() *schema.Resource {
	return define.Resources["Bridge"]
}

func (p *ListBridgeParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["list"]
}

func (p *ListBridgeParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *ListBridgeParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *ListBridgeParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *ListBridgeParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *ListBridgeParam) GetOutputFormat() string {
	return "table"
}

func (p *ListBridgeParam) SetName(v []string) {
	p.Name = v
}

func (p *ListBridgeParam) GetName() []string {
	return p.Name
}
func (p *ListBridgeParam) SetId(v []int64) {
	p.Id = v
}

func (p *ListBridgeParam) GetId() []int64 {
	return p.Id
}
func (p *ListBridgeParam) SetFrom(v int) {
	p.From = v
}

func (p *ListBridgeParam) GetFrom() int {
	return p.From
}
func (p *ListBridgeParam) SetMax(v int) {
	p.Max = v
}

func (p *ListBridgeParam) GetMax() int {
	return p.Max
}
func (p *ListBridgeParam) SetSort(v []string) {
	p.Sort = v
}

func (p *ListBridgeParam) GetSort() []string {
	return p.Sort
}
func (p *ListBridgeParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *ListBridgeParam) GetOutputType() string {
	return p.OutputType
}
func (p *ListBridgeParam) SetColumn(v []string) {
	p.Column = v
}

func (p *ListBridgeParam) GetColumn() []string {
	return p.Column
}
func (p *ListBridgeParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *ListBridgeParam) GetQuiet() bool {
	return p.Quiet
}
func (p *ListBridgeParam) SetFormat(v string) {
	p.Format = v
}

func (p *ListBridgeParam) GetFormat() string {
	return p.Format
}
func (p *ListBridgeParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *ListBridgeParam) GetFormatFile() string {
	return p.FormatFile
}

// CreateBridgeParam is input parameters for the sacloud API
type CreateBridgeParam struct {
	Name        string
	Description string
	Assumeyes   bool
	OutputType  string
	Column      []string
	Quiet       bool
	Format      string
	FormatFile  string
}

// NewCreateBridgeParam return new CreateBridgeParam
func NewCreateBridgeParam() *CreateBridgeParam {
	return &CreateBridgeParam{}
}

// Validate checks current values in model
func (p *CreateBridgeParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Bridge"].Commands["create"].Params["name"].ValidateFunc
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Bridge"].Commands["create"].Params["description"].ValidateFunc
		errs := validator("--description", p.Description)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := schema.ValidateInStrValues("json", "csv", "tsv")
		errs := validator("--output-type", p.OutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateOutputOption(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *CreateBridgeParam) GetResourceDef() *schema.Resource {
	return define.Resources["Bridge"]
}

func (p *CreateBridgeParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["create"]
}

func (p *CreateBridgeParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *CreateBridgeParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *CreateBridgeParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *CreateBridgeParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *CreateBridgeParam) GetOutputFormat() string {
	return "table"
}

func (p *CreateBridgeParam) SetName(v string) {
	p.Name = v
}

func (p *CreateBridgeParam) GetName() string {
	return p.Name
}
func (p *CreateBridgeParam) SetDescription(v string) {
	p.Description = v
}

func (p *CreateBridgeParam) GetDescription() string {
	return p.Description
}
func (p *CreateBridgeParam) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *CreateBridgeParam) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *CreateBridgeParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *CreateBridgeParam) GetOutputType() string {
	return p.OutputType
}
func (p *CreateBridgeParam) SetColumn(v []string) {
	p.Column = v
}

func (p *CreateBridgeParam) GetColumn() []string {
	return p.Column
}
func (p *CreateBridgeParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *CreateBridgeParam) GetQuiet() bool {
	return p.Quiet
}
func (p *CreateBridgeParam) SetFormat(v string) {
	p.Format = v
}

func (p *CreateBridgeParam) GetFormat() string {
	return p.Format
}
func (p *CreateBridgeParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *CreateBridgeParam) GetFormatFile() string {
	return p.FormatFile
}

// ReadBridgeParam is input parameters for the sacloud API
type ReadBridgeParam struct {
	OutputType string
	Column     []string
	Quiet      bool
	Format     string
	FormatFile string
	Id         int64
}

// NewReadBridgeParam return new ReadBridgeParam
func NewReadBridgeParam() *ReadBridgeParam {
	return &ReadBridgeParam{}
}

// Validate checks current values in model
func (p *ReadBridgeParam) Validate() []error {
	errors := []error{}
	{
		validator := validateSakuraID
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := schema.ValidateInStrValues("json", "csv", "tsv")
		errs := validator("--output-type", p.OutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateOutputOption(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *ReadBridgeParam) GetResourceDef() *schema.Resource {
	return define.Resources["Bridge"]
}

func (p *ReadBridgeParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["read"]
}

func (p *ReadBridgeParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *ReadBridgeParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *ReadBridgeParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *ReadBridgeParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *ReadBridgeParam) GetOutputFormat() string {
	return "table"
}

func (p *ReadBridgeParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *ReadBridgeParam) GetOutputType() string {
	return p.OutputType
}
func (p *ReadBridgeParam) SetColumn(v []string) {
	p.Column = v
}

func (p *ReadBridgeParam) GetColumn() []string {
	return p.Column
}
func (p *ReadBridgeParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *ReadBridgeParam) GetQuiet() bool {
	return p.Quiet
}
func (p *ReadBridgeParam) SetFormat(v string) {
	p.Format = v
}

func (p *ReadBridgeParam) GetFormat() string {
	return p.Format
}
func (p *ReadBridgeParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *ReadBridgeParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *ReadBridgeParam) SetId(v int64) {
	p.Id = v
}

func (p *ReadBridgeParam) GetId() int64 {
	return p.Id
}

// UpdateBridgeParam is input parameters for the sacloud API
type UpdateBridgeParam struct {
	Name        string
	Description string
	Assumeyes   bool
	OutputType  string
	Column      []string
	Quiet       bool
	Format      string
	FormatFile  string
	Id          int64
}

// NewUpdateBridgeParam return new UpdateBridgeParam
func NewUpdateBridgeParam() *UpdateBridgeParam {
	return &UpdateBridgeParam{}
}

// Validate checks current values in model
func (p *UpdateBridgeParam) Validate() []error {
	errors := []error{}
	{
		validator := define.Resources["Bridge"].Commands["update"].Params["name"].ValidateFunc
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Bridge"].Commands["update"].Params["description"].ValidateFunc
		errs := validator("--description", p.Description)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := validateSakuraID
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := schema.ValidateInStrValues("json", "csv", "tsv")
		errs := validator("--output-type", p.OutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateOutputOption(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *UpdateBridgeParam) GetResourceDef() *schema.Resource {
	return define.Resources["Bridge"]
}

func (p *UpdateBridgeParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["update"]
}

func (p *UpdateBridgeParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *UpdateBridgeParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *UpdateBridgeParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *UpdateBridgeParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *UpdateBridgeParam) GetOutputFormat() string {
	return "table"
}

func (p *UpdateBridgeParam) SetName(v string) {
	p.Name = v
}

func (p *UpdateBridgeParam) GetName() string {
	return p.Name
}
func (p *UpdateBridgeParam) SetDescription(v string) {
	p.Description = v
}

func (p *UpdateBridgeParam) GetDescription() string {
	return p.Description
}
func (p *UpdateBridgeParam) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *UpdateBridgeParam) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *UpdateBridgeParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *UpdateBridgeParam) GetOutputType() string {
	return p.OutputType
}
func (p *UpdateBridgeParam) SetColumn(v []string) {
	p.Column = v
}

func (p *UpdateBridgeParam) GetColumn() []string {
	return p.Column
}
func (p *UpdateBridgeParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *UpdateBridgeParam) GetQuiet() bool {
	return p.Quiet
}
func (p *UpdateBridgeParam) SetFormat(v string) {
	p.Format = v
}

func (p *UpdateBridgeParam) GetFormat() string {
	return p.Format
}
func (p *UpdateBridgeParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *UpdateBridgeParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *UpdateBridgeParam) SetId(v int64) {
	p.Id = v
}

func (p *UpdateBridgeParam) GetId() int64 {
	return p.Id
}

// DeleteBridgeParam is input parameters for the sacloud API
type DeleteBridgeParam struct {
	Assumeyes  bool
	OutputType string
	Column     []string
	Quiet      bool
	Format     string
	FormatFile string
	Id         int64
}

// NewDeleteBridgeParam return new DeleteBridgeParam
func NewDeleteBridgeParam() *DeleteBridgeParam {
	return &DeleteBridgeParam{}
}

// Validate checks current values in model
func (p *DeleteBridgeParam) Validate() []error {
	errors := []error{}
	{
		validator := validateSakuraID
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := schema.ValidateInStrValues("json", "csv", "tsv")
		errs := validator("--output-type", p.OutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateOutputOption(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *DeleteBridgeParam) GetResourceDef() *schema.Resource {
	return define.Resources["Bridge"]
}

func (p *DeleteBridgeParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["delete"]
}

func (p *DeleteBridgeParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *DeleteBridgeParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *DeleteBridgeParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *DeleteBridgeParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *DeleteBridgeParam) GetOutputFormat() string {
	return "table"
}

func (p *DeleteBridgeParam) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *DeleteBridgeParam) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *DeleteBridgeParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *DeleteBridgeParam) GetOutputType() string {
	return p.OutputType
}
func (p *DeleteBridgeParam) SetColumn(v []string) {
	p.Column = v
}

func (p *DeleteBridgeParam) GetColumn() []string {
	return p.Column
}
func (p *DeleteBridgeParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *DeleteBridgeParam) GetQuiet() bool {
	return p.Quiet
}
func (p *DeleteBridgeParam) SetFormat(v string) {
	p.Format = v
}

func (p *DeleteBridgeParam) GetFormat() string {
	return p.Format
}
func (p *DeleteBridgeParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *DeleteBridgeParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *DeleteBridgeParam) SetId(v int64) {
	p.Id = v
}

func (p *DeleteBridgeParam) GetId() int64 {
	return p.Id
}
