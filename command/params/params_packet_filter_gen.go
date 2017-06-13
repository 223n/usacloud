// Code generated by 'github.com/sacloud/usacloud/tools/gen-input-models'; DO NOT EDIT

package params

import (
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

// ListPacketFilterParam is input parameters for the sacloud API
type ListPacketFilterParam struct {
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

// NewListPacketFilterParam return new ListPacketFilterParam
func NewListPacketFilterParam() *ListPacketFilterParam {
	return &ListPacketFilterParam{}
}

// Validate checks current values in model
func (p *ListPacketFilterParam) Validate() []error {
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
		validator := define.Resources["PacketFilter"].Commands["list"].Params["id"].ValidateFunc
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

func (p *ListPacketFilterParam) GetResourceDef() *schema.Resource {
	return define.Resources["PacketFilter"]
}

func (p *ListPacketFilterParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["list"]
}

func (p *ListPacketFilterParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *ListPacketFilterParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *ListPacketFilterParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *ListPacketFilterParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *ListPacketFilterParam) GetOutputFormat() string {
	return "table"
}

func (p *ListPacketFilterParam) SetName(v []string) {
	p.Name = v
}

func (p *ListPacketFilterParam) GetName() []string {
	return p.Name
}
func (p *ListPacketFilterParam) SetId(v []int64) {
	p.Id = v
}

func (p *ListPacketFilterParam) GetId() []int64 {
	return p.Id
}
func (p *ListPacketFilterParam) SetFrom(v int) {
	p.From = v
}

func (p *ListPacketFilterParam) GetFrom() int {
	return p.From
}
func (p *ListPacketFilterParam) SetMax(v int) {
	p.Max = v
}

func (p *ListPacketFilterParam) GetMax() int {
	return p.Max
}
func (p *ListPacketFilterParam) SetSort(v []string) {
	p.Sort = v
}

func (p *ListPacketFilterParam) GetSort() []string {
	return p.Sort
}
func (p *ListPacketFilterParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *ListPacketFilterParam) GetOutputType() string {
	return p.OutputType
}
func (p *ListPacketFilterParam) SetColumn(v []string) {
	p.Column = v
}

func (p *ListPacketFilterParam) GetColumn() []string {
	return p.Column
}
func (p *ListPacketFilterParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *ListPacketFilterParam) GetQuiet() bool {
	return p.Quiet
}
func (p *ListPacketFilterParam) SetFormat(v string) {
	p.Format = v
}

func (p *ListPacketFilterParam) GetFormat() string {
	return p.Format
}
func (p *ListPacketFilterParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *ListPacketFilterParam) GetFormatFile() string {
	return p.FormatFile
}

// CreatePacketFilterParam is input parameters for the sacloud API
type CreatePacketFilterParam struct {
	Name        string
	Description string
	Assumeyes   bool
	OutputType  string
	Column      []string
	Quiet       bool
	Format      string
	FormatFile  string
}

// NewCreatePacketFilterParam return new CreatePacketFilterParam
func NewCreatePacketFilterParam() *CreatePacketFilterParam {
	return &CreatePacketFilterParam{}
}

// Validate checks current values in model
func (p *CreatePacketFilterParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["PacketFilter"].Commands["create"].Params["name"].ValidateFunc
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["PacketFilter"].Commands["create"].Params["description"].ValidateFunc
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

func (p *CreatePacketFilterParam) GetResourceDef() *schema.Resource {
	return define.Resources["PacketFilter"]
}

func (p *CreatePacketFilterParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["create"]
}

func (p *CreatePacketFilterParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *CreatePacketFilterParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *CreatePacketFilterParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *CreatePacketFilterParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *CreatePacketFilterParam) GetOutputFormat() string {
	return "table"
}

func (p *CreatePacketFilterParam) SetName(v string) {
	p.Name = v
}

func (p *CreatePacketFilterParam) GetName() string {
	return p.Name
}
func (p *CreatePacketFilterParam) SetDescription(v string) {
	p.Description = v
}

func (p *CreatePacketFilterParam) GetDescription() string {
	return p.Description
}
func (p *CreatePacketFilterParam) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *CreatePacketFilterParam) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *CreatePacketFilterParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *CreatePacketFilterParam) GetOutputType() string {
	return p.OutputType
}
func (p *CreatePacketFilterParam) SetColumn(v []string) {
	p.Column = v
}

func (p *CreatePacketFilterParam) GetColumn() []string {
	return p.Column
}
func (p *CreatePacketFilterParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *CreatePacketFilterParam) GetQuiet() bool {
	return p.Quiet
}
func (p *CreatePacketFilterParam) SetFormat(v string) {
	p.Format = v
}

func (p *CreatePacketFilterParam) GetFormat() string {
	return p.Format
}
func (p *CreatePacketFilterParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *CreatePacketFilterParam) GetFormatFile() string {
	return p.FormatFile
}

// ReadPacketFilterParam is input parameters for the sacloud API
type ReadPacketFilterParam struct {
	OutputType string
	Column     []string
	Quiet      bool
	Format     string
	FormatFile string
	Id         int64
}

// NewReadPacketFilterParam return new ReadPacketFilterParam
func NewReadPacketFilterParam() *ReadPacketFilterParam {
	return &ReadPacketFilterParam{}
}

// Validate checks current values in model
func (p *ReadPacketFilterParam) Validate() []error {
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

func (p *ReadPacketFilterParam) GetResourceDef() *schema.Resource {
	return define.Resources["PacketFilter"]
}

func (p *ReadPacketFilterParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["read"]
}

func (p *ReadPacketFilterParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *ReadPacketFilterParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *ReadPacketFilterParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *ReadPacketFilterParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *ReadPacketFilterParam) GetOutputFormat() string {
	return "table"
}

func (p *ReadPacketFilterParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *ReadPacketFilterParam) GetOutputType() string {
	return p.OutputType
}
func (p *ReadPacketFilterParam) SetColumn(v []string) {
	p.Column = v
}

func (p *ReadPacketFilterParam) GetColumn() []string {
	return p.Column
}
func (p *ReadPacketFilterParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *ReadPacketFilterParam) GetQuiet() bool {
	return p.Quiet
}
func (p *ReadPacketFilterParam) SetFormat(v string) {
	p.Format = v
}

func (p *ReadPacketFilterParam) GetFormat() string {
	return p.Format
}
func (p *ReadPacketFilterParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *ReadPacketFilterParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *ReadPacketFilterParam) SetId(v int64) {
	p.Id = v
}

func (p *ReadPacketFilterParam) GetId() int64 {
	return p.Id
}

// UpdatePacketFilterParam is input parameters for the sacloud API
type UpdatePacketFilterParam struct {
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

// NewUpdatePacketFilterParam return new UpdatePacketFilterParam
func NewUpdatePacketFilterParam() *UpdatePacketFilterParam {
	return &UpdatePacketFilterParam{}
}

// Validate checks current values in model
func (p *UpdatePacketFilterParam) Validate() []error {
	errors := []error{}
	{
		validator := define.Resources["PacketFilter"].Commands["update"].Params["name"].ValidateFunc
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["PacketFilter"].Commands["update"].Params["description"].ValidateFunc
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

func (p *UpdatePacketFilterParam) GetResourceDef() *schema.Resource {
	return define.Resources["PacketFilter"]
}

func (p *UpdatePacketFilterParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["update"]
}

func (p *UpdatePacketFilterParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *UpdatePacketFilterParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *UpdatePacketFilterParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *UpdatePacketFilterParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *UpdatePacketFilterParam) GetOutputFormat() string {
	return "table"
}

func (p *UpdatePacketFilterParam) SetName(v string) {
	p.Name = v
}

func (p *UpdatePacketFilterParam) GetName() string {
	return p.Name
}
func (p *UpdatePacketFilterParam) SetDescription(v string) {
	p.Description = v
}

func (p *UpdatePacketFilterParam) GetDescription() string {
	return p.Description
}
func (p *UpdatePacketFilterParam) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *UpdatePacketFilterParam) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *UpdatePacketFilterParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *UpdatePacketFilterParam) GetOutputType() string {
	return p.OutputType
}
func (p *UpdatePacketFilterParam) SetColumn(v []string) {
	p.Column = v
}

func (p *UpdatePacketFilterParam) GetColumn() []string {
	return p.Column
}
func (p *UpdatePacketFilterParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *UpdatePacketFilterParam) GetQuiet() bool {
	return p.Quiet
}
func (p *UpdatePacketFilterParam) SetFormat(v string) {
	p.Format = v
}

func (p *UpdatePacketFilterParam) GetFormat() string {
	return p.Format
}
func (p *UpdatePacketFilterParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *UpdatePacketFilterParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *UpdatePacketFilterParam) SetId(v int64) {
	p.Id = v
}

func (p *UpdatePacketFilterParam) GetId() int64 {
	return p.Id
}

// DeletePacketFilterParam is input parameters for the sacloud API
type DeletePacketFilterParam struct {
	Assumeyes  bool
	OutputType string
	Column     []string
	Quiet      bool
	Format     string
	FormatFile string
	Id         int64
}

// NewDeletePacketFilterParam return new DeletePacketFilterParam
func NewDeletePacketFilterParam() *DeletePacketFilterParam {
	return &DeletePacketFilterParam{}
}

// Validate checks current values in model
func (p *DeletePacketFilterParam) Validate() []error {
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

func (p *DeletePacketFilterParam) GetResourceDef() *schema.Resource {
	return define.Resources["PacketFilter"]
}

func (p *DeletePacketFilterParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["delete"]
}

func (p *DeletePacketFilterParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *DeletePacketFilterParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *DeletePacketFilterParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *DeletePacketFilterParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *DeletePacketFilterParam) GetOutputFormat() string {
	return "table"
}

func (p *DeletePacketFilterParam) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *DeletePacketFilterParam) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *DeletePacketFilterParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *DeletePacketFilterParam) GetOutputType() string {
	return p.OutputType
}
func (p *DeletePacketFilterParam) SetColumn(v []string) {
	p.Column = v
}

func (p *DeletePacketFilterParam) GetColumn() []string {
	return p.Column
}
func (p *DeletePacketFilterParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *DeletePacketFilterParam) GetQuiet() bool {
	return p.Quiet
}
func (p *DeletePacketFilterParam) SetFormat(v string) {
	p.Format = v
}

func (p *DeletePacketFilterParam) GetFormat() string {
	return p.Format
}
func (p *DeletePacketFilterParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *DeletePacketFilterParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *DeletePacketFilterParam) SetId(v int64) {
	p.Id = v
}

func (p *DeletePacketFilterParam) GetId() int64 {
	return p.Id
}

// RuleInfoPacketFilterParam is input parameters for the sacloud API
type RuleInfoPacketFilterParam struct {
	OutputType string
	Column     []string
	Quiet      bool
	Format     string
	FormatFile string
	Id         int64
}

// NewRuleInfoPacketFilterParam return new RuleInfoPacketFilterParam
func NewRuleInfoPacketFilterParam() *RuleInfoPacketFilterParam {
	return &RuleInfoPacketFilterParam{}
}

// Validate checks current values in model
func (p *RuleInfoPacketFilterParam) Validate() []error {
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

func (p *RuleInfoPacketFilterParam) GetResourceDef() *schema.Resource {
	return define.Resources["PacketFilter"]
}

func (p *RuleInfoPacketFilterParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["rule-info"]
}

func (p *RuleInfoPacketFilterParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *RuleInfoPacketFilterParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *RuleInfoPacketFilterParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *RuleInfoPacketFilterParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *RuleInfoPacketFilterParam) GetOutputFormat() string {
	return "table"
}

func (p *RuleInfoPacketFilterParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *RuleInfoPacketFilterParam) GetOutputType() string {
	return p.OutputType
}
func (p *RuleInfoPacketFilterParam) SetColumn(v []string) {
	p.Column = v
}

func (p *RuleInfoPacketFilterParam) GetColumn() []string {
	return p.Column
}
func (p *RuleInfoPacketFilterParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *RuleInfoPacketFilterParam) GetQuiet() bool {
	return p.Quiet
}
func (p *RuleInfoPacketFilterParam) SetFormat(v string) {
	p.Format = v
}

func (p *RuleInfoPacketFilterParam) GetFormat() string {
	return p.Format
}
func (p *RuleInfoPacketFilterParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *RuleInfoPacketFilterParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *RuleInfoPacketFilterParam) SetId(v int64) {
	p.Id = v
}

func (p *RuleInfoPacketFilterParam) GetId() int64 {
	return p.Id
}

// RuleAddPacketFilterParam is input parameters for the sacloud API
type RuleAddPacketFilterParam struct {
	Index           int
	Protocol        string
	SourceNetwork   string
	SourcePort      string
	DestinationPort string
	Action          string
	Description     string
	Assumeyes       bool
	OutputType      string
	Column          []string
	Quiet           bool
	Format          string
	FormatFile      string
	Id              int64
}

// NewRuleAddPacketFilterParam return new RuleAddPacketFilterParam
func NewRuleAddPacketFilterParam() *RuleAddPacketFilterParam {
	return &RuleAddPacketFilterParam{

		Index: 1,
	}
}

// Validate checks current values in model
func (p *RuleAddPacketFilterParam) Validate() []error {
	errors := []error{}
	{
		validator := define.Resources["PacketFilter"].Commands["rule-add"].Params["protocol"].ValidateFunc
		errs := validator("--protocol", p.Protocol)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["PacketFilter"].Commands["rule-add"].Params["source-network"].ValidateFunc
		errs := validator("--source-network", p.SourceNetwork)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["PacketFilter"].Commands["rule-add"].Params["source-port"].ValidateFunc
		errs := validator("--source-port", p.SourcePort)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["PacketFilter"].Commands["rule-add"].Params["destination-port"].ValidateFunc
		errs := validator("--destination-port", p.DestinationPort)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["PacketFilter"].Commands["rule-add"].Params["action"].ValidateFunc
		errs := validator("--action", p.Action)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["PacketFilter"].Commands["rule-add"].Params["description"].ValidateFunc
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

func (p *RuleAddPacketFilterParam) GetResourceDef() *schema.Resource {
	return define.Resources["PacketFilter"]
}

func (p *RuleAddPacketFilterParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["rule-add"]
}

func (p *RuleAddPacketFilterParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *RuleAddPacketFilterParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *RuleAddPacketFilterParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *RuleAddPacketFilterParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *RuleAddPacketFilterParam) GetOutputFormat() string {
	return "table"
}

func (p *RuleAddPacketFilterParam) SetIndex(v int) {
	p.Index = v
}

func (p *RuleAddPacketFilterParam) GetIndex() int {
	return p.Index
}
func (p *RuleAddPacketFilterParam) SetProtocol(v string) {
	p.Protocol = v
}

func (p *RuleAddPacketFilterParam) GetProtocol() string {
	return p.Protocol
}
func (p *RuleAddPacketFilterParam) SetSourceNetwork(v string) {
	p.SourceNetwork = v
}

func (p *RuleAddPacketFilterParam) GetSourceNetwork() string {
	return p.SourceNetwork
}
func (p *RuleAddPacketFilterParam) SetSourcePort(v string) {
	p.SourcePort = v
}

func (p *RuleAddPacketFilterParam) GetSourcePort() string {
	return p.SourcePort
}
func (p *RuleAddPacketFilterParam) SetDestinationPort(v string) {
	p.DestinationPort = v
}

func (p *RuleAddPacketFilterParam) GetDestinationPort() string {
	return p.DestinationPort
}
func (p *RuleAddPacketFilterParam) SetAction(v string) {
	p.Action = v
}

func (p *RuleAddPacketFilterParam) GetAction() string {
	return p.Action
}
func (p *RuleAddPacketFilterParam) SetDescription(v string) {
	p.Description = v
}

func (p *RuleAddPacketFilterParam) GetDescription() string {
	return p.Description
}
func (p *RuleAddPacketFilterParam) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *RuleAddPacketFilterParam) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *RuleAddPacketFilterParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *RuleAddPacketFilterParam) GetOutputType() string {
	return p.OutputType
}
func (p *RuleAddPacketFilterParam) SetColumn(v []string) {
	p.Column = v
}

func (p *RuleAddPacketFilterParam) GetColumn() []string {
	return p.Column
}
func (p *RuleAddPacketFilterParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *RuleAddPacketFilterParam) GetQuiet() bool {
	return p.Quiet
}
func (p *RuleAddPacketFilterParam) SetFormat(v string) {
	p.Format = v
}

func (p *RuleAddPacketFilterParam) GetFormat() string {
	return p.Format
}
func (p *RuleAddPacketFilterParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *RuleAddPacketFilterParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *RuleAddPacketFilterParam) SetId(v int64) {
	p.Id = v
}

func (p *RuleAddPacketFilterParam) GetId() int64 {
	return p.Id
}

// RuleUpdatePacketFilterParam is input parameters for the sacloud API
type RuleUpdatePacketFilterParam struct {
	Index           int
	Protocol        string
	SourceNetwork   string
	SourcePort      string
	DestinationPort string
	Action          string
	Description     string
	Assumeyes       bool
	OutputType      string
	Column          []string
	Quiet           bool
	Format          string
	FormatFile      string
	Id              int64
}

// NewRuleUpdatePacketFilterParam return new RuleUpdatePacketFilterParam
func NewRuleUpdatePacketFilterParam() *RuleUpdatePacketFilterParam {
	return &RuleUpdatePacketFilterParam{}
}

// Validate checks current values in model
func (p *RuleUpdatePacketFilterParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--index", p.Index)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["PacketFilter"].Commands["rule-update"].Params["protocol"].ValidateFunc
		errs := validator("--protocol", p.Protocol)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["PacketFilter"].Commands["rule-update"].Params["source-network"].ValidateFunc
		errs := validator("--source-network", p.SourceNetwork)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["PacketFilter"].Commands["rule-update"].Params["source-port"].ValidateFunc
		errs := validator("--source-port", p.SourcePort)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["PacketFilter"].Commands["rule-update"].Params["destination-port"].ValidateFunc
		errs := validator("--destination-port", p.DestinationPort)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["PacketFilter"].Commands["rule-update"].Params["action"].ValidateFunc
		errs := validator("--action", p.Action)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["PacketFilter"].Commands["rule-update"].Params["description"].ValidateFunc
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

func (p *RuleUpdatePacketFilterParam) GetResourceDef() *schema.Resource {
	return define.Resources["PacketFilter"]
}

func (p *RuleUpdatePacketFilterParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["rule-update"]
}

func (p *RuleUpdatePacketFilterParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *RuleUpdatePacketFilterParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *RuleUpdatePacketFilterParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *RuleUpdatePacketFilterParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *RuleUpdatePacketFilterParam) GetOutputFormat() string {
	return "table"
}

func (p *RuleUpdatePacketFilterParam) SetIndex(v int) {
	p.Index = v
}

func (p *RuleUpdatePacketFilterParam) GetIndex() int {
	return p.Index
}
func (p *RuleUpdatePacketFilterParam) SetProtocol(v string) {
	p.Protocol = v
}

func (p *RuleUpdatePacketFilterParam) GetProtocol() string {
	return p.Protocol
}
func (p *RuleUpdatePacketFilterParam) SetSourceNetwork(v string) {
	p.SourceNetwork = v
}

func (p *RuleUpdatePacketFilterParam) GetSourceNetwork() string {
	return p.SourceNetwork
}
func (p *RuleUpdatePacketFilterParam) SetSourcePort(v string) {
	p.SourcePort = v
}

func (p *RuleUpdatePacketFilterParam) GetSourcePort() string {
	return p.SourcePort
}
func (p *RuleUpdatePacketFilterParam) SetDestinationPort(v string) {
	p.DestinationPort = v
}

func (p *RuleUpdatePacketFilterParam) GetDestinationPort() string {
	return p.DestinationPort
}
func (p *RuleUpdatePacketFilterParam) SetAction(v string) {
	p.Action = v
}

func (p *RuleUpdatePacketFilterParam) GetAction() string {
	return p.Action
}
func (p *RuleUpdatePacketFilterParam) SetDescription(v string) {
	p.Description = v
}

func (p *RuleUpdatePacketFilterParam) GetDescription() string {
	return p.Description
}
func (p *RuleUpdatePacketFilterParam) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *RuleUpdatePacketFilterParam) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *RuleUpdatePacketFilterParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *RuleUpdatePacketFilterParam) GetOutputType() string {
	return p.OutputType
}
func (p *RuleUpdatePacketFilterParam) SetColumn(v []string) {
	p.Column = v
}

func (p *RuleUpdatePacketFilterParam) GetColumn() []string {
	return p.Column
}
func (p *RuleUpdatePacketFilterParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *RuleUpdatePacketFilterParam) GetQuiet() bool {
	return p.Quiet
}
func (p *RuleUpdatePacketFilterParam) SetFormat(v string) {
	p.Format = v
}

func (p *RuleUpdatePacketFilterParam) GetFormat() string {
	return p.Format
}
func (p *RuleUpdatePacketFilterParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *RuleUpdatePacketFilterParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *RuleUpdatePacketFilterParam) SetId(v int64) {
	p.Id = v
}

func (p *RuleUpdatePacketFilterParam) GetId() int64 {
	return p.Id
}

// RuleDeletePacketFilterParam is input parameters for the sacloud API
type RuleDeletePacketFilterParam struct {
	Index      int
	Assumeyes  bool
	OutputType string
	Column     []string
	Quiet      bool
	Format     string
	FormatFile string
	Id         int64
}

// NewRuleDeletePacketFilterParam return new RuleDeletePacketFilterParam
func NewRuleDeletePacketFilterParam() *RuleDeletePacketFilterParam {
	return &RuleDeletePacketFilterParam{}
}

// Validate checks current values in model
func (p *RuleDeletePacketFilterParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--index", p.Index)
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

func (p *RuleDeletePacketFilterParam) GetResourceDef() *schema.Resource {
	return define.Resources["PacketFilter"]
}

func (p *RuleDeletePacketFilterParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["rule-delete"]
}

func (p *RuleDeletePacketFilterParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *RuleDeletePacketFilterParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *RuleDeletePacketFilterParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *RuleDeletePacketFilterParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *RuleDeletePacketFilterParam) GetOutputFormat() string {
	return "table"
}

func (p *RuleDeletePacketFilterParam) SetIndex(v int) {
	p.Index = v
}

func (p *RuleDeletePacketFilterParam) GetIndex() int {
	return p.Index
}
func (p *RuleDeletePacketFilterParam) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *RuleDeletePacketFilterParam) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *RuleDeletePacketFilterParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *RuleDeletePacketFilterParam) GetOutputType() string {
	return p.OutputType
}
func (p *RuleDeletePacketFilterParam) SetColumn(v []string) {
	p.Column = v
}

func (p *RuleDeletePacketFilterParam) GetColumn() []string {
	return p.Column
}
func (p *RuleDeletePacketFilterParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *RuleDeletePacketFilterParam) GetQuiet() bool {
	return p.Quiet
}
func (p *RuleDeletePacketFilterParam) SetFormat(v string) {
	p.Format = v
}

func (p *RuleDeletePacketFilterParam) GetFormat() string {
	return p.Format
}
func (p *RuleDeletePacketFilterParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *RuleDeletePacketFilterParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *RuleDeletePacketFilterParam) SetId(v int64) {
	p.Id = v
}

func (p *RuleDeletePacketFilterParam) GetId() int64 {
	return p.Id
}

// InterfaceConnectPacketFilterParam is input parameters for the sacloud API
type InterfaceConnectPacketFilterParam struct {
	InterfaceId int64
	Assumeyes   bool
	Id          int64
}

// NewInterfaceConnectPacketFilterParam return new InterfaceConnectPacketFilterParam
func NewInterfaceConnectPacketFilterParam() *InterfaceConnectPacketFilterParam {
	return &InterfaceConnectPacketFilterParam{}
}

// Validate checks current values in model
func (p *InterfaceConnectPacketFilterParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--interface-id", p.InterfaceId)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["PacketFilter"].Commands["interface-connect"].Params["interface-id"].ValidateFunc
		errs := validator("--interface-id", p.InterfaceId)
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

	return errors
}

func (p *InterfaceConnectPacketFilterParam) GetResourceDef() *schema.Resource {
	return define.Resources["PacketFilter"]
}

func (p *InterfaceConnectPacketFilterParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["interface-connect"]
}

func (p *InterfaceConnectPacketFilterParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *InterfaceConnectPacketFilterParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *InterfaceConnectPacketFilterParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *InterfaceConnectPacketFilterParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *InterfaceConnectPacketFilterParam) GetOutputFormat() string {
	return "table"
}

func (p *InterfaceConnectPacketFilterParam) SetInterfaceId(v int64) {
	p.InterfaceId = v
}

func (p *InterfaceConnectPacketFilterParam) GetInterfaceId() int64 {
	return p.InterfaceId
}
func (p *InterfaceConnectPacketFilterParam) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *InterfaceConnectPacketFilterParam) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *InterfaceConnectPacketFilterParam) SetId(v int64) {
	p.Id = v
}

func (p *InterfaceConnectPacketFilterParam) GetId() int64 {
	return p.Id
}

// InterfaceDisconnectPacketFilterParam is input parameters for the sacloud API
type InterfaceDisconnectPacketFilterParam struct {
	InterfaceId int64
	Assumeyes   bool
	Id          int64
}

// NewInterfaceDisconnectPacketFilterParam return new InterfaceDisconnectPacketFilterParam
func NewInterfaceDisconnectPacketFilterParam() *InterfaceDisconnectPacketFilterParam {
	return &InterfaceDisconnectPacketFilterParam{}
}

// Validate checks current values in model
func (p *InterfaceDisconnectPacketFilterParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--interface-id", p.InterfaceId)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["PacketFilter"].Commands["interface-disconnect"].Params["interface-id"].ValidateFunc
		errs := validator("--interface-id", p.InterfaceId)
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

	return errors
}

func (p *InterfaceDisconnectPacketFilterParam) GetResourceDef() *schema.Resource {
	return define.Resources["PacketFilter"]
}

func (p *InterfaceDisconnectPacketFilterParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["interface-disconnect"]
}

func (p *InterfaceDisconnectPacketFilterParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *InterfaceDisconnectPacketFilterParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *InterfaceDisconnectPacketFilterParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *InterfaceDisconnectPacketFilterParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *InterfaceDisconnectPacketFilterParam) GetOutputFormat() string {
	return "table"
}

func (p *InterfaceDisconnectPacketFilterParam) SetInterfaceId(v int64) {
	p.InterfaceId = v
}

func (p *InterfaceDisconnectPacketFilterParam) GetInterfaceId() int64 {
	return p.InterfaceId
}
func (p *InterfaceDisconnectPacketFilterParam) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *InterfaceDisconnectPacketFilterParam) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *InterfaceDisconnectPacketFilterParam) SetId(v int64) {
	p.Id = v
}

func (p *InterfaceDisconnectPacketFilterParam) GetId() int64 {
	return p.Id
}
