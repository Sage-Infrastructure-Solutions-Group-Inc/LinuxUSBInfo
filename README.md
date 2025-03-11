# LinuxUSBInfo
This is a very simple program that hooks into libudev (need the development libs installed to compile) and serializes 
the properties of connected devices to a JSON list. 
## Example Output
```JSON
[{"BUSNUM":"001","CURRENT_TAGS":":seat:","DEVNAME":"/dev/bus/usb/001/001","DEVNUM":"001","DEVPATH":"/devices/platform/axi/1000120000.pcie/1f00200000.usb/xhci-hcd.0/usb1","DEVTYPE":"usb_device","DRIVER":"usb","ID_AUTOSUSPEND":"1","ID_BUS":"usb","ID_FOR_SEAT":"usb-platform-xhci-hcd_0","ID_MODEL":"xHCI_Host_Controller","ID_MODEL_ENC":"xHCI\\x20Host\\x20Controller","ID_MODEL_FROM_DATABASE":"2.0 root hub","ID_MODEL_ID":"0002","ID_PATH":"platform-xhci-hcd.0","ID_PATH_TAG":"platform-xhci-hcd_0","ID_REVISION":"0606","ID_SERIAL":"Linux_6.6.51+rpt-rpi-2712_xhci-hcd_xHCI_Host_Controller_xhci-hcd.0","ID_SERIAL_SHORT":"xhci-hcd.0","ID_USB_INTERFACES":":090000:","ID_USB_MODEL":"xHCI_Host_Controller","ID_USB_MODEL_ENC":"xHCI\\x20Host\\x20Controller","ID_USB_MODEL_ID":"0002","ID_USB_REVISION":"0606","ID_USB_SERIAL":"Linux_6.6.51+rpt-rpi-2712_xhci-hcd_xHCI_Host_Controller_xhci-hcd.0","ID_USB_SERIAL_SHORT":"xhci-hcd.0","ID_USB_VENDOR":"Linux_6.6.51+rpt-rpi-2712_xhci-hcd","ID_USB_VENDOR_ENC":"Linux\\x206.6.51+rpt-rpi-2712\\x20xhci-hcd","ID_USB_VENDOR_ID":"1d6b","ID_VENDOR":"Linux_6.6.51+rpt-rpi-2712_xhci-hcd","ID_VENDOR_ENC":"Linux\\x206.6.51+rpt-rpi-2712\\x20xhci-hcd","ID_VENDOR_FROM_DATABASE":"Linux Foundation","ID_VENDOR_ID":"1d6b","MAJOR":"189","MINOR":"0","OF_ALIAS_0":"usb0","OF_COMPATIBLE_0":"snps,dwc3","OF_COMPATIBLE_N":"1","OF_FULLNAME":"/axi/pcie@120000/rp1/usb@200000","OF_NAME":"usb","PRODUCT":"1d6b/2/606","SUBSYSTEM":"usb","TAGS":":seat:","TYPE":"9/0/1","USEC_INITIALIZED":"975275"},{"DEVPATH":"/devices/platform/axi/1000120000.pcie/1f00200000.usb/xhci-hcd.0/usb1/1-0:1.0","DEVTYPE":"usb_interface","DRIVER":"hub","ID_AUTOSUSPEND":"1","ID_MODEL_FROM_DATABASE":"2.0 root hub","ID_PATH":"platform-xhci-hcd.0-usb-0:0:1.0","ID_PATH_TAG":"platform-xhci-hcd_0-usb-0_0_1_0","ID_USB_CLASS_FROM_DATABASE":"Hub","ID_USB_PROTOCOL_FROM_DATABASE":"Single TT","ID_VENDOR_FROM_DATABASE":"Linux Foundation","INTERFACE":"9/0/0","MODALIAS":"usb:v1D6Bp0002d0606dc09dsc00dp01ic09isc00ip00in00","OF_ALIAS_0":"usb0","OF_COMPATIBLE_0":"snps,dwc3","OF_COMPATIBLE_N":"1","OF_FULLNAME":"/axi/pcie@120000/rp1/usb@200000","OF_NAME":"usb","PRODUCT":"1d6b/2/606","SUBSYSTEM":"usb","TYPE":"9/0/1","USEC_INITIALIZED":"982614"},{"BUSNUM":"002","CURRENT_TAGS":":seat:","DEVNAME":"/dev/bus/usb/002/001","DEVNUM":"001","DEVPATH":"/devices/platform/axi/1000120000.pcie/1f00200000.usb/xhci-hcd.0/usb2","DEVTYPE":"usb_device","DRIVER":"usb","ID_AUTOSUSPEND":"1","ID_BUS":"usb","ID_FOR_SEAT":"usb-platform-xhci-hcd_0","ID_MODEL":"xHCI_Host_Controller","ID_MODEL_ENC":"xHCI\\x20Host\\x20Controller","ID_MODEL_FROM_DATABASE":"3.0 root hub","ID_MODEL_ID":"0003","ID_PATH":"platform-xhci-hcd.0","ID_PATH_TAG":"platform-xhci-hcd_0","ID_REVISION":"0606","ID_SERIAL":"Linux_6.6.51+rpt-rpi-2712_xhci-hcd_xHCI_Host_Controller_xhci-hcd.0","ID_SERIAL_SHORT":"xhci-hcd.0","ID_USB_INTERFACES":":090000:","ID_USB_MODEL":"xHCI_Host_Controller","ID_USB_MODEL_ENC":"xHCI\\x20Host\\x20Controller","ID_USB_MODEL_ID":"0003","ID_USB_REVISION":"0606","ID_USB_SERIAL":"Linux_6.6.51+rpt-rpi-2712_xhci-hcd_xHCI_Host_Controller_xhci-hcd.0","ID_USB_SERIAL_SHORT":"xhci-hcd.0","ID_USB_VENDOR":"Linux_6.6.51+rpt-rpi-2712_xhci-hcd","ID_USB_VENDOR_ENC":"Linux\\x206.6.51+rpt-rpi-2712\\x20xhci-hcd","ID_USB_VENDOR_ID":"1d6b","ID_VENDOR":"Linux_6.6.51+rpt-rpi-2712_xhci-hcd","ID_VENDOR_ENC":"Linux\\x206.6.51+rpt-rpi-2712\\x20xhci-hcd","ID_VENDOR_FROM_DATABASE":"Linux Foundation","ID_VENDOR_ID":"1d6b","MAJOR":"189","MINOR":"128","OF_ALIAS_0":"usb0","OF_COMPATIBLE_0":"snps,dwc3","OF_COMPATIBLE_N":"1","OF_FULLNAME":"/axi/pcie@120000/rp1/usb@200000","OF_NAME":"usb","PRODUCT":"1d6b/3/606","SUBSYSTEM":"usb","TAGS":":seat:","TYPE":"9/0/3","USEC_INITIALIZED":"953016"},{"DEVPATH":"/devices/platform/axi/1000120000.pcie/1f00200000.usb/xhci-hcd.0/usb2/2-0:1.0","DEVTYPE":"usb_interface","DRIVER":"hub","ID_AUTOSUSPEND":"1","ID_MODEL_FROM_DATABASE":"3.0 root hub","ID_PATH":"platform-xhci-hcd.0-usb-0:0:1.0","ID_PATH_TAG":"platform-xhci-hcd_0-usb-0_0_1_0","ID_USB_CLASS_FROM_DATABASE":"Hub","ID_VENDOR_FROM_DATABASE":"Linux Foundation","INTERFACE":"9/0/0","MODALIAS":"usb:v1D6Bp0003d0606dc09dsc00dp03ic09isc00ip00in00","OF_ALIAS_0":"usb0","OF_COMPATIBLE_0":"snps,dwc3","OF_COMPATIBLE_N":"1","OF_FULLNAME":"/axi/pcie@120000/rp1/usb@200000","OF_NAME":"usb","PRODUCT":"1d6b/3/606","SUBSYSTEM":"usb","TYPE":"9/0/3","USEC_INITIALIZED":"981033"},{"BUSNUM":"003","CURRENT_TAGS":":seat:","DEVNAME":"/dev/bus/usb/003/001","DEVNUM":"001","DEVPATH":"/devices/platform/axi/1000120000.pcie/1f00300000.usb/xhci-hcd.1/usb3","DEVTYPE":"usb_device","DRIVER":"usb","ID_AUTOSUSPEND":"1","ID_BUS":"usb","ID_FOR_SEAT":"usb-platform-xhci-hcd_1","ID_MODEL":"xHCI_Host_Controller","ID_MODEL_ENC":"xHCI\\x20Host\\x20Controller","ID_MODEL_FROM_DATABASE":"2.0 root hub","ID_MODEL_ID":"0002","ID_PATH":"platform-xhci-hcd.1","ID_PATH_TAG":"platform-xhci-hcd_1","ID_REVISION":"0606","ID_SERIAL":"Linux_6.6.51+rpt-rpi-2712_xhci-hcd_xHCI_Host_Controller_xhci-hcd.1","ID_SERIAL_SHORT":"xhci-hcd.1","ID_USB_INTERFACES":":090000:","ID_USB_MODEL":"xHCI_Host_Controller","ID_USB_MODEL_ENC":"xHCI\\x20Host\\x20Controller","ID_USB_MODEL_ID":"0002","ID_USB_REVISION":"0606","ID_USB_SERIAL":"Linux_6.6.51+rpt-rpi-2712_xhci-hcd_xHCI_Host_Controller_xhci-hcd.1","ID_USB_SERIAL_SHORT":"xhci-hcd.1","ID_USB_VENDOR":"Linux_6.6.51+rpt-rpi-2712_xhci-hcd","ID_USB_VENDOR_ENC":"Linux\\x206.6.51+rpt-rpi-2712\\x20xhci-hcd","ID_USB_VENDOR_ID":"1d6b","ID_VENDOR":"Linux_6.6.51+rpt-rpi-2712_xhci-hcd","ID_VENDOR_ENC":"Linux\\x206.6.51+rpt-rpi-2712\\x20xhci-hcd","ID_VENDOR_FROM_DATABASE":"Linux Foundation","ID_VENDOR_ID":"1d6b","MAJOR":"189","MINOR":"256","OF_ALIAS_0":"usb1","OF_COMPATIBLE_0":"snps,dwc3","OF_COMPATIBLE_N":"1","OF_FULLNAME":"/axi/pcie@120000/rp1/usb@300000","OF_NAME":"usb","PRODUCT":"1d6b/2/606","SUBSYSTEM":"usb","TAGS":":seat:","TYPE":"9/0/1","USEC_INITIALIZED":"985908"},{"DEVPATH":"/devices/platform/axi/1000120000.pcie/1f00300000.usb/xhci-hcd.1/usb3/3-0:1.0","DEVTYPE":"usb_interface","DRIVER":"hub","ID_AUTOSUSPEND":"1","ID_MODEL_FROM_DATABASE":"2.0 root hub","ID_PATH":"platform-xhci-hcd.1-usb-0:0:1.0","ID_PATH_TAG":"platform-xhci-hcd_1-usb-0_0_1_0","ID_USB_CLASS_FROM_DATABASE":"Hub","ID_USB_PROTOCOL_FROM_DATABASE":"Single TT","ID_VENDOR_FROM_DATABASE":"Linux Foundation","INTERFACE":"9/0/0","MODALIAS":"usb:v1D6Bp0002d0606dc09dsc00dp01ic09isc00ip00in00","OF_ALIAS_0":"usb1","OF_COMPATIBLE_0":"snps,dwc3","OF_COMPATIBLE_N":"1","OF_FULLNAME":"/axi/pcie@120000/rp1/usb@300000","OF_NAME":"usb","PRODUCT":"1d6b/2/606","SUBSYSTEM":"usb","TYPE":"9/0/1","USEC_INITIALIZED":"999880"},{"BUSNUM":"003","CURRENT_TAGS":":seat:","DEVNAME":"/dev/bus/usb/003/002","DEVNUM":"002","DEVPATH":"/devices/platform/axi/1000120000.pcie/1f00300000.usb/xhci-hcd.1/usb3/3-2","DEVTYPE":"usb_device","DRIVER":"usb","ID_BUS":"usb","ID_DRIVE_THUMB":"1","ID_FOR_SEAT":"usb-platform-xhci-hcd_1-usb-0_2","ID_MODEL":"DataTraveler_3.0","ID_MODEL_ENC":"DataTraveler\\x203.0","ID_MODEL_FROM_DATABASE":"DataTraveler 100 G3/G4/SE9 G2/50 Kyson","ID_MODEL_ID":"1666","ID_PATH":"platform-xhci-hcd.1-usb-0:2","ID_PATH_TAG":"platform-xhci-hcd_1-usb-0_2","ID_REVISION":"0200","ID_SERIAL":"Kingston_DataTraveler_3.0_FA2A448DB7E716A1B1880920","ID_SERIAL_SHORT":"FA2A448DB7E716A1B1880920","ID_USB_INTERFACES":":080650:","ID_USB_MODEL":"DataTraveler_3.0","ID_USB_MODEL_ENC":"DataTraveler\\x203.0","ID_USB_MODEL_ID":"1666","ID_USB_REVISION":"0200","ID_USB_SERIAL":"Kingston_DataTraveler_3.0_FA2A448DB7E716A1B1880920","ID_USB_SERIAL_SHORT":"FA2A448DB7E716A1B1880920","ID_USB_VENDOR":"Kingston","ID_USB_VENDOR_ENC":"Kingston","ID_USB_VENDOR_ID":"0951","ID_VENDOR":"Kingston","ID_VENDOR_ENC":"Kingston","ID_VENDOR_FROM_DATABASE":"Kingston Technology","ID_VENDOR_ID":"0951","MAJOR":"189","MINOR":"257","PRODUCT":"951/1666/200","SUBSYSTEM":"usb","TAGS":":seat:","TYPE":"0/0/0","USEC_INITIALIZED":"259714520"},{"DEVPATH":"/devices/platform/axi/1000120000.pcie/1f00300000.usb/xhci-hcd.1/usb3/3-2/3-2:1.0","DEVTYPE":"usb_interface","DRIVER":"usb-storage","ID_MODEL_FROM_DATABASE":"DataTraveler 100 G3/G4/SE9 G2/50 Kyson","ID_PATH":"platform-xhci-hcd.1-usb-0:2:1.0","ID_PATH_TAG":"platform-xhci-hcd_1-usb-0_2_1_0","ID_VENDOR_FROM_DATABASE":"Kingston Technology","INTERFACE":"8/6/80","MODALIAS":"usb:v0951p1666d0200dc00dsc00dp00ic08isc06ip50in00","PRODUCT":"951/1666/200","SUBSYSTEM":"usb","TYPE":"0/0/0","USEC_INITIALIZED":"259716091"},{"BUSNUM":"004","CURRENT_TAGS":":seat:","DEVNAME":"/dev/bus/usb/004/001","DEVNUM":"001","DEVPATH":"/devices/platform/axi/1000120000.pcie/1f00300000.usb/xhci-hcd.1/usb4","DEVTYPE":"usb_device","DRIVER":"usb","ID_AUTOSUSPEND":"1","ID_BUS":"usb","ID_FOR_SEAT":"usb-platform-xhci-hcd_1","ID_MODEL":"xHCI_Host_Controller","ID_MODEL_ENC":"xHCI\\x20Host\\x20Controller","ID_MODEL_FROM_DATABASE":"3.0 root hub","ID_MODEL_ID":"0003","ID_PATH":"platform-xhci-hcd.1","ID_PATH_TAG":"platform-xhci-hcd_1","ID_REVISION":"0606","ID_SERIAL":"Linux_6.6.51+rpt-rpi-2712_xhci-hcd_xHCI_Host_Controller_xhci-hcd.1","ID_SERIAL_SHORT":"xhci-hcd.1","ID_USB_INTERFACES":":090000:","ID_USB_MODEL":"xHCI_Host_Controller","ID_USB_MODEL_ENC":"xHCI\\x20Host\\x20Controller","ID_USB_MODEL_ID":"0003","ID_USB_REVISION":"0606","ID_USB_SERIAL":"Linux_6.6.51+rpt-rpi-2712_xhci-hcd_xHCI_Host_Controller_xhci-hcd.1","ID_USB_SERIAL_SHORT":"xhci-hcd.1","ID_USB_VENDOR":"Linux_6.6.51+rpt-rpi-2712_xhci-hcd","ID_USB_VENDOR_ENC":"Linux\\x206.6.51+rpt-rpi-2712\\x20xhci-hcd","ID_USB_VENDOR_ID":"1d6b","ID_VENDOR":"Linux_6.6.51+rpt-rpi-2712_xhci-hcd","ID_VENDOR_ENC":"Linux\\x206.6.51+rpt-rpi-2712\\x20xhci-hcd","ID_VENDOR_FROM_DATABASE":"Linux Foundation","ID_VENDOR_ID":"1d6b","MAJOR":"189","MINOR":"384","OF_ALIAS_0":"usb1","OF_COMPATIBLE_0":"snps,dwc3","OF_COMPATIBLE_N":"1","OF_FULLNAME":"/axi/pcie@120000/rp1/usb@300000","OF_NAME":"usb","PRODUCT":"1d6b/3/606","SUBSYSTEM":"usb","TAGS":":seat:","TYPE":"9/0/3","USEC_INITIALIZED":"938826"},{"DEVPATH":"/devices/platform/axi/1000120000.pcie/1f00300000.usb/xhci-hcd.1/usb4/4-0:1.0","DEVTYPE":"usb_interface","DRIVER":"hub","ID_AUTOSUSPEND":"1","ID_MODEL_FROM_DATABASE":"3.0 root hub","ID_PATH":"platform-xhci-hcd.1-usb-0:0:1.0","ID_PATH_TAG":"platform-xhci-hcd_1-usb-0_0_1_0","ID_USB_CLASS_FROM_DATABASE":"Hub","ID_VENDOR_FROM_DATABASE":"Linux Foundation","INTERFACE":"9/0/0","MODALIAS":"usb:v1D6Bp0003d0606dc09dsc00dp03ic09isc00ip00in00","OF_ALIAS_0":"usb1","OF_COMPATIBLE_0":"snps,dwc3","OF_COMPATIBLE_N":"1","OF_FULLNAME":"/axi/pcie@120000/rp1/usb@300000","OF_NAME":"usb","PRODUCT":"1d6b/3/606","SUBSYSTEM":"usb","TYPE":"9/0/3","USEC_INITIALIZED":"951845"}]
```
> No newline is printed to hopefully make this data a little easier to use.

# Compilation
To compile you need to:
1. Clone this repository (`git clone https://github.com/Sage-Infrastructure-Solutions-Group-Inc/LinuxUSBInfo.git`)
2. Install GoLang
3. Install libudev (On Debian the package is `libudev-dev`)
4. Build or run the software
5. Enjoy!