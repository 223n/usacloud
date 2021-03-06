// Code generated by 'github.com/sacloud/usacloud/tools/gen-input-models'; DO NOT EDIT

package params

import (
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

// ListSimpleMonitorParam is input parameters for the sacloud API
type ListSimpleMonitorParam struct {
	Name              []string `json:"name"`
	Id                []int64  `json:"id"`
	Tags              []string `json:"tags"`
	From              int      `json:"from"`
	Max               int      `json:"max"`
	Sort              []string `json:"sort"`
	ParamTemplate     string   `json:"param-template"`
	ParamTemplateFile string   `json:"param-template-file"`
	GenerateSkeleton  bool     `json:"generate-skeleton"`
	OutputType        string   `json:"output-type"`
	Column            []string `json:"column"`
	Quiet             bool     `json:"quiet"`
	Format            string   `json:"format"`
	FormatFile        string   `json:"format-file"`
	Query             string   `json:"query"`
}

// NewListSimpleMonitorParam return new ListSimpleMonitorParam
func NewListSimpleMonitorParam() *ListSimpleMonitorParam {
	return &ListSimpleMonitorParam{}
}

// FillValueToSkeleton fill values to empty fields
func (p *ListSimpleMonitorParam) FillValueToSkeleton() {
	if isEmpty(p.Name) {
		p.Name = []string{""}
	}
	if isEmpty(p.Id) {
		p.Id = []int64{0}
	}
	if isEmpty(p.Tags) {
		p.Tags = []string{""}
	}
	if isEmpty(p.From) {
		p.From = 0
	}
	if isEmpty(p.Max) {
		p.Max = 0
	}
	if isEmpty(p.Sort) {
		p.Sort = []string{""}
	}
	if isEmpty(p.ParamTemplate) {
		p.ParamTemplate = ""
	}
	if isEmpty(p.ParamTemplateFile) {
		p.ParamTemplateFile = ""
	}
	if isEmpty(p.GenerateSkeleton) {
		p.GenerateSkeleton = false
	}
	if isEmpty(p.OutputType) {
		p.OutputType = ""
	}
	if isEmpty(p.Column) {
		p.Column = []string{""}
	}
	if isEmpty(p.Quiet) {
		p.Quiet = false
	}
	if isEmpty(p.Format) {
		p.Format = ""
	}
	if isEmpty(p.FormatFile) {
		p.FormatFile = ""
	}
	if isEmpty(p.Query) {
		p.Query = ""
	}

}

// Validate checks current values in model
func (p *ListSimpleMonitorParam) Validate() []error {
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
		validator := define.Resources["SimpleMonitor"].Commands["list"].Params["id"].ValidateFunc
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
		validator := define.Resources["SimpleMonitor"].Commands["list"].Params["tags"].ValidateFunc
		errs := validator("--tags", p.Tags)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := schema.ValidateInStrValues(define.AllowOutputTypes...)
		errs := validator("--output-type", p.OutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateInputOption(p)
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

func (p *ListSimpleMonitorParam) GetResourceDef() *schema.Resource {
	return define.Resources["SimpleMonitor"]
}

func (p *ListSimpleMonitorParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["list"]
}

func (p *ListSimpleMonitorParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *ListSimpleMonitorParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *ListSimpleMonitorParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *ListSimpleMonitorParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *ListSimpleMonitorParam) SetName(v []string) {
	p.Name = v
}

func (p *ListSimpleMonitorParam) GetName() []string {
	return p.Name
}
func (p *ListSimpleMonitorParam) SetId(v []int64) {
	p.Id = v
}

func (p *ListSimpleMonitorParam) GetId() []int64 {
	return p.Id
}
func (p *ListSimpleMonitorParam) SetTags(v []string) {
	p.Tags = v
}

func (p *ListSimpleMonitorParam) GetTags() []string {
	return p.Tags
}
func (p *ListSimpleMonitorParam) SetFrom(v int) {
	p.From = v
}

func (p *ListSimpleMonitorParam) GetFrom() int {
	return p.From
}
func (p *ListSimpleMonitorParam) SetMax(v int) {
	p.Max = v
}

func (p *ListSimpleMonitorParam) GetMax() int {
	return p.Max
}
func (p *ListSimpleMonitorParam) SetSort(v []string) {
	p.Sort = v
}

func (p *ListSimpleMonitorParam) GetSort() []string {
	return p.Sort
}
func (p *ListSimpleMonitorParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *ListSimpleMonitorParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *ListSimpleMonitorParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *ListSimpleMonitorParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *ListSimpleMonitorParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *ListSimpleMonitorParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *ListSimpleMonitorParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *ListSimpleMonitorParam) GetOutputType() string {
	return p.OutputType
}
func (p *ListSimpleMonitorParam) SetColumn(v []string) {
	p.Column = v
}

func (p *ListSimpleMonitorParam) GetColumn() []string {
	return p.Column
}
func (p *ListSimpleMonitorParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *ListSimpleMonitorParam) GetQuiet() bool {
	return p.Quiet
}
func (p *ListSimpleMonitorParam) SetFormat(v string) {
	p.Format = v
}

func (p *ListSimpleMonitorParam) GetFormat() string {
	return p.Format
}
func (p *ListSimpleMonitorParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *ListSimpleMonitorParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *ListSimpleMonitorParam) SetQuery(v string) {
	p.Query = v
}

func (p *ListSimpleMonitorParam) GetQuery() string {
	return p.Query
}

// CreateSimpleMonitorParam is input parameters for the sacloud API
type CreateSimpleMonitorParam struct {
	Target            string   `json:"target"`
	Protocol          string   `json:"protocol"`
	Port              int      `json:"port"`
	DelayLoop         int      `json:"delay-loop"`
	Disabled          bool     `json:"disabled"`
	HostHeader        string   `json:"host-header"`
	Path              string   `json:"path"`
	ResponseCode      int      `json:"response-code"`
	Sni               bool     `json:"sni"`
	DnsQname          string   `json:"dns-qname"`
	DnsExcepted       string   `json:"dns-excepted"`
	RemainingDays     int      `json:"remaining-days"`
	NotifyEmail       bool     `json:"notify-email"`
	EmailType         string   `json:"email-type"`
	SlackWebhook      string   `json:"slack-webhook"`
	Description       string   `json:"description"`
	Tags              []string `json:"tags"`
	IconId            int64    `json:"icon-id"`
	Assumeyes         bool     `json:"assumeyes"`
	ParamTemplate     string   `json:"param-template"`
	ParamTemplateFile string   `json:"param-template-file"`
	GenerateSkeleton  bool     `json:"generate-skeleton"`
	OutputType        string   `json:"output-type"`
	Column            []string `json:"column"`
	Quiet             bool     `json:"quiet"`
	Format            string   `json:"format"`
	FormatFile        string   `json:"format-file"`
	Query             string   `json:"query"`
}

// NewCreateSimpleMonitorParam return new CreateSimpleMonitorParam
func NewCreateSimpleMonitorParam() *CreateSimpleMonitorParam {
	return &CreateSimpleMonitorParam{

		Protocol:      "ping",
		DelayLoop:     1,
		RemainingDays: 30,
		NotifyEmail:   true,
		EmailType:     "text",
	}
}

// FillValueToSkeleton fill values to empty fields
func (p *CreateSimpleMonitorParam) FillValueToSkeleton() {
	if isEmpty(p.Target) {
		p.Target = ""
	}
	if isEmpty(p.Protocol) {
		p.Protocol = ""
	}
	if isEmpty(p.Port) {
		p.Port = 0
	}
	if isEmpty(p.DelayLoop) {
		p.DelayLoop = 0
	}
	if isEmpty(p.Disabled) {
		p.Disabled = false
	}
	if isEmpty(p.HostHeader) {
		p.HostHeader = ""
	}
	if isEmpty(p.Path) {
		p.Path = ""
	}
	if isEmpty(p.ResponseCode) {
		p.ResponseCode = 0
	}
	if isEmpty(p.Sni) {
		p.Sni = false
	}
	if isEmpty(p.DnsQname) {
		p.DnsQname = ""
	}
	if isEmpty(p.DnsExcepted) {
		p.DnsExcepted = ""
	}
	if isEmpty(p.RemainingDays) {
		p.RemainingDays = 0
	}
	if isEmpty(p.NotifyEmail) {
		p.NotifyEmail = false
	}
	if isEmpty(p.EmailType) {
		p.EmailType = ""
	}
	if isEmpty(p.SlackWebhook) {
		p.SlackWebhook = ""
	}
	if isEmpty(p.Description) {
		p.Description = ""
	}
	if isEmpty(p.Tags) {
		p.Tags = []string{""}
	}
	if isEmpty(p.IconId) {
		p.IconId = 0
	}
	if isEmpty(p.Assumeyes) {
		p.Assumeyes = false
	}
	if isEmpty(p.ParamTemplate) {
		p.ParamTemplate = ""
	}
	if isEmpty(p.ParamTemplateFile) {
		p.ParamTemplateFile = ""
	}
	if isEmpty(p.GenerateSkeleton) {
		p.GenerateSkeleton = false
	}
	if isEmpty(p.OutputType) {
		p.OutputType = ""
	}
	if isEmpty(p.Column) {
		p.Column = []string{""}
	}
	if isEmpty(p.Quiet) {
		p.Quiet = false
	}
	if isEmpty(p.Format) {
		p.Format = ""
	}
	if isEmpty(p.FormatFile) {
		p.FormatFile = ""
	}
	if isEmpty(p.Query) {
		p.Query = ""
	}

}

// Validate checks current values in model
func (p *CreateSimpleMonitorParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--target", p.Target)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := validateRequired
		errs := validator("--protocol", p.Protocol)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["SimpleMonitor"].Commands["create"].Params["protocol"].ValidateFunc
		errs := validator("--protocol", p.Protocol)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["SimpleMonitor"].Commands["create"].Params["port"].ValidateFunc
		errs := validator("--port", p.Port)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := validateRequired
		errs := validator("--delay-loop", p.DelayLoop)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["SimpleMonitor"].Commands["create"].Params["delay-loop"].ValidateFunc
		errs := validator("--delay-loop", p.DelayLoop)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["SimpleMonitor"].Commands["create"].Params["remaining-days"].ValidateFunc
		errs := validator("--remaining-days", p.RemainingDays)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["SimpleMonitor"].Commands["create"].Params["email-type"].ValidateFunc
		errs := validator("--email-type", p.EmailType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["SimpleMonitor"].Commands["create"].Params["description"].ValidateFunc
		errs := validator("--description", p.Description)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["SimpleMonitor"].Commands["create"].Params["tags"].ValidateFunc
		errs := validator("--tags", p.Tags)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["SimpleMonitor"].Commands["create"].Params["icon-id"].ValidateFunc
		errs := validator("--icon-id", p.IconId)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := schema.ValidateInStrValues(define.AllowOutputTypes...)
		errs := validator("--output-type", p.OutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateInputOption(p)
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

func (p *CreateSimpleMonitorParam) GetResourceDef() *schema.Resource {
	return define.Resources["SimpleMonitor"]
}

func (p *CreateSimpleMonitorParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["create"]
}

func (p *CreateSimpleMonitorParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *CreateSimpleMonitorParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *CreateSimpleMonitorParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *CreateSimpleMonitorParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *CreateSimpleMonitorParam) SetTarget(v string) {
	p.Target = v
}

func (p *CreateSimpleMonitorParam) GetTarget() string {
	return p.Target
}
func (p *CreateSimpleMonitorParam) SetProtocol(v string) {
	p.Protocol = v
}

func (p *CreateSimpleMonitorParam) GetProtocol() string {
	return p.Protocol
}
func (p *CreateSimpleMonitorParam) SetPort(v int) {
	p.Port = v
}

func (p *CreateSimpleMonitorParam) GetPort() int {
	return p.Port
}
func (p *CreateSimpleMonitorParam) SetDelayLoop(v int) {
	p.DelayLoop = v
}

func (p *CreateSimpleMonitorParam) GetDelayLoop() int {
	return p.DelayLoop
}
func (p *CreateSimpleMonitorParam) SetDisabled(v bool) {
	p.Disabled = v
}

func (p *CreateSimpleMonitorParam) GetDisabled() bool {
	return p.Disabled
}
func (p *CreateSimpleMonitorParam) SetHostHeader(v string) {
	p.HostHeader = v
}

func (p *CreateSimpleMonitorParam) GetHostHeader() string {
	return p.HostHeader
}
func (p *CreateSimpleMonitorParam) SetPath(v string) {
	p.Path = v
}

func (p *CreateSimpleMonitorParam) GetPath() string {
	return p.Path
}
func (p *CreateSimpleMonitorParam) SetResponseCode(v int) {
	p.ResponseCode = v
}

func (p *CreateSimpleMonitorParam) GetResponseCode() int {
	return p.ResponseCode
}
func (p *CreateSimpleMonitorParam) SetSni(v bool) {
	p.Sni = v
}

func (p *CreateSimpleMonitorParam) GetSni() bool {
	return p.Sni
}
func (p *CreateSimpleMonitorParam) SetDnsQname(v string) {
	p.DnsQname = v
}

func (p *CreateSimpleMonitorParam) GetDnsQname() string {
	return p.DnsQname
}
func (p *CreateSimpleMonitorParam) SetDnsExcepted(v string) {
	p.DnsExcepted = v
}

func (p *CreateSimpleMonitorParam) GetDnsExcepted() string {
	return p.DnsExcepted
}
func (p *CreateSimpleMonitorParam) SetRemainingDays(v int) {
	p.RemainingDays = v
}

func (p *CreateSimpleMonitorParam) GetRemainingDays() int {
	return p.RemainingDays
}
func (p *CreateSimpleMonitorParam) SetNotifyEmail(v bool) {
	p.NotifyEmail = v
}

func (p *CreateSimpleMonitorParam) GetNotifyEmail() bool {
	return p.NotifyEmail
}
func (p *CreateSimpleMonitorParam) SetEmailType(v string) {
	p.EmailType = v
}

func (p *CreateSimpleMonitorParam) GetEmailType() string {
	return p.EmailType
}
func (p *CreateSimpleMonitorParam) SetSlackWebhook(v string) {
	p.SlackWebhook = v
}

func (p *CreateSimpleMonitorParam) GetSlackWebhook() string {
	return p.SlackWebhook
}
func (p *CreateSimpleMonitorParam) SetDescription(v string) {
	p.Description = v
}

func (p *CreateSimpleMonitorParam) GetDescription() string {
	return p.Description
}
func (p *CreateSimpleMonitorParam) SetTags(v []string) {
	p.Tags = v
}

func (p *CreateSimpleMonitorParam) GetTags() []string {
	return p.Tags
}
func (p *CreateSimpleMonitorParam) SetIconId(v int64) {
	p.IconId = v
}

func (p *CreateSimpleMonitorParam) GetIconId() int64 {
	return p.IconId
}
func (p *CreateSimpleMonitorParam) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *CreateSimpleMonitorParam) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *CreateSimpleMonitorParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *CreateSimpleMonitorParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *CreateSimpleMonitorParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *CreateSimpleMonitorParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *CreateSimpleMonitorParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *CreateSimpleMonitorParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *CreateSimpleMonitorParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *CreateSimpleMonitorParam) GetOutputType() string {
	return p.OutputType
}
func (p *CreateSimpleMonitorParam) SetColumn(v []string) {
	p.Column = v
}

func (p *CreateSimpleMonitorParam) GetColumn() []string {
	return p.Column
}
func (p *CreateSimpleMonitorParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *CreateSimpleMonitorParam) GetQuiet() bool {
	return p.Quiet
}
func (p *CreateSimpleMonitorParam) SetFormat(v string) {
	p.Format = v
}

func (p *CreateSimpleMonitorParam) GetFormat() string {
	return p.Format
}
func (p *CreateSimpleMonitorParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *CreateSimpleMonitorParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *CreateSimpleMonitorParam) SetQuery(v string) {
	p.Query = v
}

func (p *CreateSimpleMonitorParam) GetQuery() string {
	return p.Query
}

// ReadSimpleMonitorParam is input parameters for the sacloud API
type ReadSimpleMonitorParam struct {
	Selector          []string `json:"selector"`
	ParamTemplate     string   `json:"param-template"`
	ParamTemplateFile string   `json:"param-template-file"`
	GenerateSkeleton  bool     `json:"generate-skeleton"`
	OutputType        string   `json:"output-type"`
	Column            []string `json:"column"`
	Quiet             bool     `json:"quiet"`
	Format            string   `json:"format"`
	FormatFile        string   `json:"format-file"`
	Query             string   `json:"query"`
	Id                int64    `json:"id"`
}

// NewReadSimpleMonitorParam return new ReadSimpleMonitorParam
func NewReadSimpleMonitorParam() *ReadSimpleMonitorParam {
	return &ReadSimpleMonitorParam{}
}

// FillValueToSkeleton fill values to empty fields
func (p *ReadSimpleMonitorParam) FillValueToSkeleton() {
	if isEmpty(p.Selector) {
		p.Selector = []string{""}
	}
	if isEmpty(p.ParamTemplate) {
		p.ParamTemplate = ""
	}
	if isEmpty(p.ParamTemplateFile) {
		p.ParamTemplateFile = ""
	}
	if isEmpty(p.GenerateSkeleton) {
		p.GenerateSkeleton = false
	}
	if isEmpty(p.OutputType) {
		p.OutputType = ""
	}
	if isEmpty(p.Column) {
		p.Column = []string{""}
	}
	if isEmpty(p.Quiet) {
		p.Quiet = false
	}
	if isEmpty(p.Format) {
		p.Format = ""
	}
	if isEmpty(p.FormatFile) {
		p.FormatFile = ""
	}
	if isEmpty(p.Query) {
		p.Query = ""
	}
	if isEmpty(p.Id) {
		p.Id = 0
	}

}

// Validate checks current values in model
func (p *ReadSimpleMonitorParam) Validate() []error {
	errors := []error{}
	{
		validator := validateSakuraID
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := schema.ValidateInStrValues(define.AllowOutputTypes...)
		errs := validator("--output-type", p.OutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateInputOption(p)
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

func (p *ReadSimpleMonitorParam) GetResourceDef() *schema.Resource {
	return define.Resources["SimpleMonitor"]
}

func (p *ReadSimpleMonitorParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["read"]
}

func (p *ReadSimpleMonitorParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *ReadSimpleMonitorParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *ReadSimpleMonitorParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *ReadSimpleMonitorParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *ReadSimpleMonitorParam) SetSelector(v []string) {
	p.Selector = v
}

func (p *ReadSimpleMonitorParam) GetSelector() []string {
	return p.Selector
}
func (p *ReadSimpleMonitorParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *ReadSimpleMonitorParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *ReadSimpleMonitorParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *ReadSimpleMonitorParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *ReadSimpleMonitorParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *ReadSimpleMonitorParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *ReadSimpleMonitorParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *ReadSimpleMonitorParam) GetOutputType() string {
	return p.OutputType
}
func (p *ReadSimpleMonitorParam) SetColumn(v []string) {
	p.Column = v
}

func (p *ReadSimpleMonitorParam) GetColumn() []string {
	return p.Column
}
func (p *ReadSimpleMonitorParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *ReadSimpleMonitorParam) GetQuiet() bool {
	return p.Quiet
}
func (p *ReadSimpleMonitorParam) SetFormat(v string) {
	p.Format = v
}

func (p *ReadSimpleMonitorParam) GetFormat() string {
	return p.Format
}
func (p *ReadSimpleMonitorParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *ReadSimpleMonitorParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *ReadSimpleMonitorParam) SetQuery(v string) {
	p.Query = v
}

func (p *ReadSimpleMonitorParam) GetQuery() string {
	return p.Query
}
func (p *ReadSimpleMonitorParam) SetId(v int64) {
	p.Id = v
}

func (p *ReadSimpleMonitorParam) GetId() int64 {
	return p.Id
}

// UpdateSimpleMonitorParam is input parameters for the sacloud API
type UpdateSimpleMonitorParam struct {
	Protocol          string   `json:"protocol"`
	Port              int      `json:"port"`
	DelayLoop         int      `json:"delay-loop"`
	Disabled          bool     `json:"disabled"`
	HostHeader        string   `json:"host-header"`
	Path              string   `json:"path"`
	ResponseCode      int      `json:"response-code"`
	Sni               bool     `json:"sni"`
	DnsQname          string   `json:"dns-qname"`
	DnsExcepted       string   `json:"dns-excepted"`
	RemainingDays     int      `json:"remaining-days"`
	NotifyEmail       bool     `json:"notify-email"`
	EmailType         string   `json:"email-type"`
	SlackWebhook      string   `json:"slack-webhook"`
	Selector          []string `json:"selector"`
	Description       string   `json:"description"`
	Tags              []string `json:"tags"`
	IconId            int64    `json:"icon-id"`
	Assumeyes         bool     `json:"assumeyes"`
	ParamTemplate     string   `json:"param-template"`
	ParamTemplateFile string   `json:"param-template-file"`
	GenerateSkeleton  bool     `json:"generate-skeleton"`
	OutputType        string   `json:"output-type"`
	Column            []string `json:"column"`
	Quiet             bool     `json:"quiet"`
	Format            string   `json:"format"`
	FormatFile        string   `json:"format-file"`
	Query             string   `json:"query"`
	Id                int64    `json:"id"`
}

// NewUpdateSimpleMonitorParam return new UpdateSimpleMonitorParam
func NewUpdateSimpleMonitorParam() *UpdateSimpleMonitorParam {
	return &UpdateSimpleMonitorParam{}
}

// FillValueToSkeleton fill values to empty fields
func (p *UpdateSimpleMonitorParam) FillValueToSkeleton() {
	if isEmpty(p.Protocol) {
		p.Protocol = ""
	}
	if isEmpty(p.Port) {
		p.Port = 0
	}
	if isEmpty(p.DelayLoop) {
		p.DelayLoop = 0
	}
	if isEmpty(p.Disabled) {
		p.Disabled = false
	}
	if isEmpty(p.HostHeader) {
		p.HostHeader = ""
	}
	if isEmpty(p.Path) {
		p.Path = ""
	}
	if isEmpty(p.ResponseCode) {
		p.ResponseCode = 0
	}
	if isEmpty(p.Sni) {
		p.Sni = false
	}
	if isEmpty(p.DnsQname) {
		p.DnsQname = ""
	}
	if isEmpty(p.DnsExcepted) {
		p.DnsExcepted = ""
	}
	if isEmpty(p.RemainingDays) {
		p.RemainingDays = 0
	}
	if isEmpty(p.NotifyEmail) {
		p.NotifyEmail = false
	}
	if isEmpty(p.EmailType) {
		p.EmailType = ""
	}
	if isEmpty(p.SlackWebhook) {
		p.SlackWebhook = ""
	}
	if isEmpty(p.Selector) {
		p.Selector = []string{""}
	}
	if isEmpty(p.Description) {
		p.Description = ""
	}
	if isEmpty(p.Tags) {
		p.Tags = []string{""}
	}
	if isEmpty(p.IconId) {
		p.IconId = 0
	}
	if isEmpty(p.Assumeyes) {
		p.Assumeyes = false
	}
	if isEmpty(p.ParamTemplate) {
		p.ParamTemplate = ""
	}
	if isEmpty(p.ParamTemplateFile) {
		p.ParamTemplateFile = ""
	}
	if isEmpty(p.GenerateSkeleton) {
		p.GenerateSkeleton = false
	}
	if isEmpty(p.OutputType) {
		p.OutputType = ""
	}
	if isEmpty(p.Column) {
		p.Column = []string{""}
	}
	if isEmpty(p.Quiet) {
		p.Quiet = false
	}
	if isEmpty(p.Format) {
		p.Format = ""
	}
	if isEmpty(p.FormatFile) {
		p.FormatFile = ""
	}
	if isEmpty(p.Query) {
		p.Query = ""
	}
	if isEmpty(p.Id) {
		p.Id = 0
	}

}

// Validate checks current values in model
func (p *UpdateSimpleMonitorParam) Validate() []error {
	errors := []error{}
	{
		validator := define.Resources["SimpleMonitor"].Commands["update"].Params["protocol"].ValidateFunc
		errs := validator("--protocol", p.Protocol)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["SimpleMonitor"].Commands["update"].Params["port"].ValidateFunc
		errs := validator("--port", p.Port)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["SimpleMonitor"].Commands["update"].Params["delay-loop"].ValidateFunc
		errs := validator("--delay-loop", p.DelayLoop)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["SimpleMonitor"].Commands["update"].Params["remaining-days"].ValidateFunc
		errs := validator("--remaining-days", p.RemainingDays)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["SimpleMonitor"].Commands["update"].Params["email-type"].ValidateFunc
		errs := validator("--email-type", p.EmailType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["SimpleMonitor"].Commands["update"].Params["description"].ValidateFunc
		errs := validator("--description", p.Description)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["SimpleMonitor"].Commands["update"].Params["tags"].ValidateFunc
		errs := validator("--tags", p.Tags)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["SimpleMonitor"].Commands["update"].Params["icon-id"].ValidateFunc
		errs := validator("--icon-id", p.IconId)
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
		validator := schema.ValidateInStrValues(define.AllowOutputTypes...)
		errs := validator("--output-type", p.OutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateInputOption(p)
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

func (p *UpdateSimpleMonitorParam) GetResourceDef() *schema.Resource {
	return define.Resources["SimpleMonitor"]
}

func (p *UpdateSimpleMonitorParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["update"]
}

func (p *UpdateSimpleMonitorParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *UpdateSimpleMonitorParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *UpdateSimpleMonitorParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *UpdateSimpleMonitorParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *UpdateSimpleMonitorParam) SetProtocol(v string) {
	p.Protocol = v
}

func (p *UpdateSimpleMonitorParam) GetProtocol() string {
	return p.Protocol
}
func (p *UpdateSimpleMonitorParam) SetPort(v int) {
	p.Port = v
}

func (p *UpdateSimpleMonitorParam) GetPort() int {
	return p.Port
}
func (p *UpdateSimpleMonitorParam) SetDelayLoop(v int) {
	p.DelayLoop = v
}

func (p *UpdateSimpleMonitorParam) GetDelayLoop() int {
	return p.DelayLoop
}
func (p *UpdateSimpleMonitorParam) SetDisabled(v bool) {
	p.Disabled = v
}

func (p *UpdateSimpleMonitorParam) GetDisabled() bool {
	return p.Disabled
}
func (p *UpdateSimpleMonitorParam) SetHostHeader(v string) {
	p.HostHeader = v
}

func (p *UpdateSimpleMonitorParam) GetHostHeader() string {
	return p.HostHeader
}
func (p *UpdateSimpleMonitorParam) SetPath(v string) {
	p.Path = v
}

func (p *UpdateSimpleMonitorParam) GetPath() string {
	return p.Path
}
func (p *UpdateSimpleMonitorParam) SetResponseCode(v int) {
	p.ResponseCode = v
}

func (p *UpdateSimpleMonitorParam) GetResponseCode() int {
	return p.ResponseCode
}
func (p *UpdateSimpleMonitorParam) SetSni(v bool) {
	p.Sni = v
}

func (p *UpdateSimpleMonitorParam) GetSni() bool {
	return p.Sni
}
func (p *UpdateSimpleMonitorParam) SetDnsQname(v string) {
	p.DnsQname = v
}

func (p *UpdateSimpleMonitorParam) GetDnsQname() string {
	return p.DnsQname
}
func (p *UpdateSimpleMonitorParam) SetDnsExcepted(v string) {
	p.DnsExcepted = v
}

func (p *UpdateSimpleMonitorParam) GetDnsExcepted() string {
	return p.DnsExcepted
}
func (p *UpdateSimpleMonitorParam) SetRemainingDays(v int) {
	p.RemainingDays = v
}

func (p *UpdateSimpleMonitorParam) GetRemainingDays() int {
	return p.RemainingDays
}
func (p *UpdateSimpleMonitorParam) SetNotifyEmail(v bool) {
	p.NotifyEmail = v
}

func (p *UpdateSimpleMonitorParam) GetNotifyEmail() bool {
	return p.NotifyEmail
}
func (p *UpdateSimpleMonitorParam) SetEmailType(v string) {
	p.EmailType = v
}

func (p *UpdateSimpleMonitorParam) GetEmailType() string {
	return p.EmailType
}
func (p *UpdateSimpleMonitorParam) SetSlackWebhook(v string) {
	p.SlackWebhook = v
}

func (p *UpdateSimpleMonitorParam) GetSlackWebhook() string {
	return p.SlackWebhook
}
func (p *UpdateSimpleMonitorParam) SetSelector(v []string) {
	p.Selector = v
}

func (p *UpdateSimpleMonitorParam) GetSelector() []string {
	return p.Selector
}
func (p *UpdateSimpleMonitorParam) SetDescription(v string) {
	p.Description = v
}

func (p *UpdateSimpleMonitorParam) GetDescription() string {
	return p.Description
}
func (p *UpdateSimpleMonitorParam) SetTags(v []string) {
	p.Tags = v
}

func (p *UpdateSimpleMonitorParam) GetTags() []string {
	return p.Tags
}
func (p *UpdateSimpleMonitorParam) SetIconId(v int64) {
	p.IconId = v
}

func (p *UpdateSimpleMonitorParam) GetIconId() int64 {
	return p.IconId
}
func (p *UpdateSimpleMonitorParam) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *UpdateSimpleMonitorParam) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *UpdateSimpleMonitorParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *UpdateSimpleMonitorParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *UpdateSimpleMonitorParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *UpdateSimpleMonitorParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *UpdateSimpleMonitorParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *UpdateSimpleMonitorParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *UpdateSimpleMonitorParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *UpdateSimpleMonitorParam) GetOutputType() string {
	return p.OutputType
}
func (p *UpdateSimpleMonitorParam) SetColumn(v []string) {
	p.Column = v
}

func (p *UpdateSimpleMonitorParam) GetColumn() []string {
	return p.Column
}
func (p *UpdateSimpleMonitorParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *UpdateSimpleMonitorParam) GetQuiet() bool {
	return p.Quiet
}
func (p *UpdateSimpleMonitorParam) SetFormat(v string) {
	p.Format = v
}

func (p *UpdateSimpleMonitorParam) GetFormat() string {
	return p.Format
}
func (p *UpdateSimpleMonitorParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *UpdateSimpleMonitorParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *UpdateSimpleMonitorParam) SetQuery(v string) {
	p.Query = v
}

func (p *UpdateSimpleMonitorParam) GetQuery() string {
	return p.Query
}
func (p *UpdateSimpleMonitorParam) SetId(v int64) {
	p.Id = v
}

func (p *UpdateSimpleMonitorParam) GetId() int64 {
	return p.Id
}

// DeleteSimpleMonitorParam is input parameters for the sacloud API
type DeleteSimpleMonitorParam struct {
	Selector          []string `json:"selector"`
	Assumeyes         bool     `json:"assumeyes"`
	ParamTemplate     string   `json:"param-template"`
	ParamTemplateFile string   `json:"param-template-file"`
	GenerateSkeleton  bool     `json:"generate-skeleton"`
	OutputType        string   `json:"output-type"`
	Column            []string `json:"column"`
	Quiet             bool     `json:"quiet"`
	Format            string   `json:"format"`
	FormatFile        string   `json:"format-file"`
	Query             string   `json:"query"`
	Id                int64    `json:"id"`
}

// NewDeleteSimpleMonitorParam return new DeleteSimpleMonitorParam
func NewDeleteSimpleMonitorParam() *DeleteSimpleMonitorParam {
	return &DeleteSimpleMonitorParam{}
}

// FillValueToSkeleton fill values to empty fields
func (p *DeleteSimpleMonitorParam) FillValueToSkeleton() {
	if isEmpty(p.Selector) {
		p.Selector = []string{""}
	}
	if isEmpty(p.Assumeyes) {
		p.Assumeyes = false
	}
	if isEmpty(p.ParamTemplate) {
		p.ParamTemplate = ""
	}
	if isEmpty(p.ParamTemplateFile) {
		p.ParamTemplateFile = ""
	}
	if isEmpty(p.GenerateSkeleton) {
		p.GenerateSkeleton = false
	}
	if isEmpty(p.OutputType) {
		p.OutputType = ""
	}
	if isEmpty(p.Column) {
		p.Column = []string{""}
	}
	if isEmpty(p.Quiet) {
		p.Quiet = false
	}
	if isEmpty(p.Format) {
		p.Format = ""
	}
	if isEmpty(p.FormatFile) {
		p.FormatFile = ""
	}
	if isEmpty(p.Query) {
		p.Query = ""
	}
	if isEmpty(p.Id) {
		p.Id = 0
	}

}

// Validate checks current values in model
func (p *DeleteSimpleMonitorParam) Validate() []error {
	errors := []error{}
	{
		validator := validateSakuraID
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := schema.ValidateInStrValues(define.AllowOutputTypes...)
		errs := validator("--output-type", p.OutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateInputOption(p)
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

func (p *DeleteSimpleMonitorParam) GetResourceDef() *schema.Resource {
	return define.Resources["SimpleMonitor"]
}

func (p *DeleteSimpleMonitorParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["delete"]
}

func (p *DeleteSimpleMonitorParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *DeleteSimpleMonitorParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *DeleteSimpleMonitorParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *DeleteSimpleMonitorParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *DeleteSimpleMonitorParam) SetSelector(v []string) {
	p.Selector = v
}

func (p *DeleteSimpleMonitorParam) GetSelector() []string {
	return p.Selector
}
func (p *DeleteSimpleMonitorParam) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *DeleteSimpleMonitorParam) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *DeleteSimpleMonitorParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *DeleteSimpleMonitorParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *DeleteSimpleMonitorParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *DeleteSimpleMonitorParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *DeleteSimpleMonitorParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *DeleteSimpleMonitorParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *DeleteSimpleMonitorParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *DeleteSimpleMonitorParam) GetOutputType() string {
	return p.OutputType
}
func (p *DeleteSimpleMonitorParam) SetColumn(v []string) {
	p.Column = v
}

func (p *DeleteSimpleMonitorParam) GetColumn() []string {
	return p.Column
}
func (p *DeleteSimpleMonitorParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *DeleteSimpleMonitorParam) GetQuiet() bool {
	return p.Quiet
}
func (p *DeleteSimpleMonitorParam) SetFormat(v string) {
	p.Format = v
}

func (p *DeleteSimpleMonitorParam) GetFormat() string {
	return p.Format
}
func (p *DeleteSimpleMonitorParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *DeleteSimpleMonitorParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *DeleteSimpleMonitorParam) SetQuery(v string) {
	p.Query = v
}

func (p *DeleteSimpleMonitorParam) GetQuery() string {
	return p.Query
}
func (p *DeleteSimpleMonitorParam) SetId(v int64) {
	p.Id = v
}

func (p *DeleteSimpleMonitorParam) GetId() int64 {
	return p.Id
}
