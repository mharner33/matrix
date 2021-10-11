package matrix

//This file includes most of the key data structures used by the matrix rules engine

//List of Acclerator types supported
const (
	Gpu = iota + 1
	Fpga
	Nic
	Nvme
	Other
)

//List of Acclerator Vendors
const (
	Nvidia = iota + 1
	Mellanox
	Intel
	AMD
	Emulex
	Xilinx
	Samsung
	Toshiba
)

//Physical width of accelerators
const (
	Single = iota + 1
	Double
	Triple
)

//List of PCIe Generations
const (
	Gen3 = iota + 3
	Gen4
	Gen5
	Gen6
)
const (
	X4  = 4
	X8  = 8
	X16 = 16
)

//PCIe device type - this determines how much BAR memory is required.
const (
	Type1 = iota + 1
	Type2
	Type3
	Type4
)

//Top or bottom Director
const (
	Top = iota + 1
	Bottom
)

//List of server vendors
const (
	Dell = iota + 1
	Supermicro
	Lenovo
	HPe
	Cisco
	Generic
)

//Struct for accelerator cards
type Accelerator struct {
	Atype     int  //Acclerator type - use constants above
	Power     int  //power in Watts
	PcieGen   int  //PCIe generation - 3, 4, etc
	SlotWidth int  //physical size of accelerator
	ExtPort   bool //Is there an external port required
	PcieSpeed int
	DevType   int //PCIe device type as recognized by Liqid
	DevCnt    int //Number of PCIe device ID's to
}

//Struct for the different PCIe expansion chassis
type PcieChassis struct {
	Slots       int
	PcieGen     int
	HostPorts   int //this is the number of x4 ports available for connectivity
	FabricPorts int //Number of x4 ports used to connect to switch infrastucture
	Btu         int //Heat output of chassis
	BasePower   int
	Throughput  int           //Throughput in GB/s
	AcceList    []Accelerator //list of accelerators
	Direct      bool          //Will this be a direct connect configuration
	DevCnt      int           //Number of PCIe devices taken by this device.
	Cables      int           //Number of included PCIe cables
}

//Struct for the different PCIe switches
type PcieSwitch struct {
	Ports       int //Number of PCIe x4 ports
	PcieGen     int //PCIe Generation
	BasePower   int //Power useage of switch
	Btu         int //Heat output of switch
	HostPorts   int //Number of x4 ports connected to hosts
	FabricPorts int //Number of x4 ports connected to Chassis
	CrossPorts  int //Number of x4 ports connected to additional switch
	DevCnt      int //Number of PCIe devices taken by this switch
}

//Struct for various Host types
type Host struct {
	Vendor    int //Manufacturer of server
	Model     string
	PcieSpeed int //Speed this host will be connected - i.e., x4, x8, x16
	PcieGen   int
	DualHBA   bool
	BasePower int
	Btu       int
}

//Struct for the Liqid Director
type Director struct {
	MgmtPorts int //Number of ports available for managment
	DirType   int //Used in Hierarchical designs to determine top switch or bottom
	DevCnt    int //Number of PCIe devices taken by the director
	Cables    int //Number of included management cables
}

//Struct for the entire infrastructure
type Infrastructure struct {
	Hosts      []Host        //Slice of hosts in the Infrastructure
	Chassis    []PcieChassis //Slice of chassis in the infrastructure
	Switches   []PcieSwitch  //Slice of switches in the infrastructure
	PciCables  int           //Number of PCIe cables needed
	MgmtCables int           //Number of PCIe cables needed
	GpuCables  int           //Number of GPU Power Cables needed
}
