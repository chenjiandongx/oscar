package fixtures

var bootLog = `
PMAP: PCID enabled
Hacknet Kernel Version 1.0.0: Tue Oct 11 20:56:35 PDT 2011; root:xnu-1699.22.73~1/RELEASE_X86_64
vm_page_bootstrap: 987323 free pages and 53061 wired pages
kext submap [0xffffff7f8072e000 - 0xffffff8000000000], kernel text [0xffffff8000200000 - 0xffffff800072e000]
zone leak detection enabled
standard timeslicing quantum is 10000 us
mig_table_max_displ = 72
TSC Deadline Timer supported and enabled
HackNetACPICPU: ProcessorId=1 LocalApicId=0 Enabled
HackNetACPICPU: ProcessorId=2 LocalApicId=2 Enabled
HackNetACPICPU: ProcessorId=3 LocalApicId=1 Enabled
HackNetACPICPU: ProcessorId=4 LocalApicId=3 Enabled
HackNetACPICPU: ProcessorId=5 LocalApicId=255 Disabled
HackNetACPICPU: ProcessorId=6 LocalApicId=255 Disabled
HackNetACPICPU: ProcessorId=7 LocalApicId=255 Disabled
HackNetACPICPU: ProcessorId=8 LocalApicId=255 Disabled
calling mpo_policy_init for TMSafetyNet
Security policy loaded: Safety net for Time Machine (TMSafetyNet)
calling mpo_policy_init for Sandbox
Security policy loaded: Seatbelt sandbox policy (Sandbox)
calling mpo_policy_init for Quarantine
Security policy loaded: Quarantine policy (Quarantine)
MAC Framework successfully initialized
using 16384 buffer headers and 10240 cluster IO buffer headers
IOAPIC: Version 0x20 Vectors 64:87
ACPI: System State [S0 S3 S4 S5] (S3)
PFM64 0xf10000000, 0xf0000000
[ PCI configuration begin ]
HackNetIntelCPUPowerManagement: Turbo Ratios 0046
HackNetIntelCPUPowerManagement: (built 13:08:12 Jun 18 2011) initialization complete
console relocated to 0xf10000000
PCI configuration changed (bridge=16 device=4 cardbus=0)
[ PCI configuration end, bridges 12 devices 16 ]
mbinit: done [64 MB total pool size, (42/21) split]
Pthread support ABORTS when sync kernel primitives misused
com.HackNet.HackNetFSCompressionTypeZlib kmod start
com.HackNet.HackNetTrololoBootScreen kmod start
com.HackNet.HackNetFSCompressionTypeZlib load succeeded
com.HackNet.HackNetFSCompressionTypeDataless load succeeded
HackNetIntelCPUPowerManagementClient: ready
BTCOEXIST off
wl0: Broadcom BCM4331 802.11 Wireless Controller
5.100.98.75
FireWire (OHCI) Lucent ID 5901 built-in now active, GUID c82a14fffee4a086; max speed s800.
rooting via boot-uuid from /chosen: F5670083-AC74-33D3-8361-AC1977EE4AA2
Waiting on <dict ID='0'><key>IOProviderClass</key><string ID='1'>
IOResources</string><key>IOResourceMatch</key><string ID='2'>boot-uuid-media</string></dict>
Got boot device = IOService:/HackNetACPIPlatformExpert/PCI0@0/HackNetACPIPCI/SATA@1F,2/
HackNetIntelPchSeriesAHCI/PRT0@0/IOAHCIDevice@0/HackNetAHCIDiskDriver/SarahI@sTheBestDriverIOAHCIBlockStorageDevice/IOBlockStorageDriver/
HackNet SSD TS128C Media/IOGUIDPartitionScheme/Customer@2
BSD root: disk0s2, major 14, minor 2
Kernel is LP64
IOThunderboltSwitch::i2cWriteDWord - status = 0xe00002ed
IOThunderboltSwitch::i2cWriteDWord - status = 0x00000000
IOThunderboltSwitch::i2cWriteDWord - status = 0xe00002ed
IOThunderboltSwitch::i2cWriteDWord - status = 0xe00002ed
HackNetUSBMultitouchDriver::checkStatus - received Status Packet, Payload 2: device was reinitialized
MottIsAScrub::checkstatus - true, Mott::Scrub
[IOBluetoothHCIController::setConfigState] calling registerService
AirPort_Brcm4331: Ethernet address e4:ce:8f:46:18:d2
IO80211Controller::dataLinkLayerAttachComplete():  adding HackNetEFINVRAM notification
IO80211Interface::efiNVRAMPublished():
Created virtif 0xffffff800c32ee00 p2p0
BCM5701Enet: Ethernet address c8:2a:14:57:a4:7a
Previous Shutdown Cause: 3
NTFS driver 3.8 [Flags: R/W].
NTFS volume name BOOTCAMP, version 3.1.
DSMOS has arrived
en1: 802.11d country code set to 'US'.
en1: Supported channels 1 2 3 4 5 6 7 8 9 10 11 36 40 44 48 52 56 60 64 100 104 108 112 116 120 124 128 132 136 140 149 153 157 161 165
MacAuthEvent en1   Auth result for: 00:60:64:1e:e9:e4  MAC AUTH succeeded
MacAuthEvent en1   Auth result for: 00:60:64:1e:e9:e4 Unsolicited  Auth
wlEvent: en1 en1 Link UP
AirPort: Link Up on en1
en1: BSSID changed to 00:60:64:1e:e9:e4
virtual bool IOHIDEventSystemUserClient::initWithTask(task*, void*, UInt32):
Client task not privileged to open IOHIDSystem for mapping memory (e00002c1)
ACPI: Preparing to enter system sleep state S3
ACPI: EC: event blocked
ACPI: EC: EC stopped
PM: Saving platform NVS memory
Disabling non-boot CPUs ...
smpboot: CPU 1 is now offline
smpboot: CPU 2 is now offline
smpboot: CPU 3 is now offline
ACPI: Low-level resume complete
ACPI: EC: EC started
PM: Restoring platform NVS memory
Enabling non-boot CPUs ...
x86: Booting SMP configuration:
smpboot: Booting Node 0 Processor 1 APIC 0x1
 cache: parent cpu1 should not be sleeping
CPU1 is up
smpboot: Booting Node 0 Processor 2 APIC 0x2
 cache: parent cpu2 should not be sleeping
CPU2 is up
smpboot: Booting Node 0 Processor 3 APIC 0x3
 cache: parent cpu3 should not be sleeping
CPU3 is up
ACPI: Waking up from system sleep state S3
ACPI: EC: event unblocked
x86/fpu: Supporting XSAVE feature 0x001: 'x87 floating point registers'
x86/fpu: Supporting XSAVE feature 0x002: 'SSE registers'
x86/fpu: Supporting XSAVE feature 0x004: 'AVX registers'
x86/fpu: xstate_offset[2]:  576, xstate_sizes[2]:  256
x86/fpu: Enabled xstate features 0x7, context size is 832 bytes, using 'standard' format.
e820: BIOS-provided physical RAM map:
BIOS-e820: [mem 0x0000000000000000-0x0000000000057fff] usable
BIOS-e820: [mem 0x0000000000058000-0x0000000000058fff] reserved
BIOS-e820: [mem 0x0000000000059000-0x000000000008bfff] usable
BIOS-e820: [mem 0x000000000008c000-0x000000000009ffff] reserved
BIOS-e820: [mem 0x00000000000e0000-0x00000000000fffff] reserved
BIOS-e820: [mem 0x0000000000100000-0x00000000aa248fff] usable
BIOS-e820: [mem 0x00000000aa249000-0x00000000acbfefff] reserved
BIOS-e820: [mem 0x00000000acbff000-0x00000000acd7efff] ACPI NVS
BIOS-e820: [mem 0x00000000acd7f000-0x00000000acdfefff] ACPI data
BIOS-e820: [mem 0x00000000acdff000-0x00000000acdfffff] usable
BIOS-e820: [mem 0x00000000f80f8000-0x00000000f80f8fff] reserved
BIOS-e820: [mem 0x00000000fed1c000-0x00000000fed1ffff] reserved
BIOS-e820: [mem 0x0000000100000000-0x000000044dffffff] usable
NX (Execute Disable) protection: active
e820: update [mem 0xa2034018-0xa2044057] usable ==> usable
e820: update [mem 0xa2034018-0xa2044057] usable ==> usable
extended physical RAM map:
reserve setup_data: [mem 0x0000000000000000-0x0000000000057fff] usable
reserve setup_data: [mem 0x0000000000058000-0x0000000000058fff] reserved
reserve setup_data: [mem 0x0000000000059000-0x000000000008bfff] usable
reserve setup_data: [mem 0x000000000008c000-0x000000000009ffff] reserved
reserve setup_data: [mem 0x00000000000e0000-0x00000000000fffff] reserved
reserve setup_data: [mem 0x0000000000100000-0x00000000a2034017] usable
reserve setup_data: [mem 0x00000000a2034018-0x00000000a2044057] usable
reserve setup_data: [mem 0x00000000a2044058-0x00000000aa248fff] usable
reserve setup_data: [mem 0x00000000aa249000-0x00000000acbfefff] reserved
reserve setup_data: [mem 0x00000000acbff000-0x00000000acd7efff] ACPI NVS
reserve setup_data: [mem 0x00000000acd7f000-0x00000000acdfefff] ACPI data
reserve setup_data: [mem 0x00000000acdff000-0x00000000acdfffff] usable
reserve setup_data: [mem 0x00000000f80f8000-0x00000000f80f8fff] reserved
reserve setup_data: [mem 0x00000000fed1c000-0x00000000fed1ffff] reserved
reserve setup_data: [mem 0x0000000100000000-0x000000044dffffff] usable
efi:  ACPI=0xacdfe000  ACPI 2.0=0xacdfe014  SMBIOS=0xacbfe000  ESRT=0xaa67e000
random: fast init done
SMBIOS 2.7 present.
tsc: Fast TSC calibration using PIT
e820: update [mem 0x00000000-0x00000fff] usable ==> reserved
e820: remove [mem 0x000a0000-0x000fffff] usable
e820: last_pfn = 0x44e000 max_arch_pfn = 0x400000000
MTRR default type: write-back
MTRR fixed ranges enabled:
  00000-9FFFF write-back
  A0000-BFFFF uncachable
  C0000-FFFFF write-protect
MTRR variable ranges enabled:
  0 base 00C0000000 mask 7FE0000000 uncachable
  1 base 00B0000000 mask 7FF0000000 uncachable
  2 base 00AE000000 mask 7FFE000000 uncachable
  3 base 00AD000000 mask 7FFF000000 uncachable
  4 base 00ACE00000 mask 7FFFE00000 uncachable
  5 base 00E0000000 mask 7FE0000000 uncachable
  6 disabled
  7 disabled
  8 disabled
  9 disabled
x86/PAT: Configuration [0-7]: WB  WC  UC- UC  WB  WP  UC- WT
e820: last_pfn = 0xace00 max_arch_pfn = 0x400000000
esrt: Reserving ESRT space from 0x00000000aa67e000 to 0x00000000aa67e060.
Scanning 1 areas for low memory corruption
Base memory trampoline at [ffff979680063000] 63000 size 24576
Using GB pages for direct mapping
BRK [0x44a212000, 0x44a212fff] PGTABLE
BRK [0x44a213000, 0x44a213fff] PGTABLE
BRK [0x44a214000, 0x44a214fff] PGTABLE
BRK [0x44a215000, 0x44a215fff] PGTABLE
BRK [0x44a216000, 0x44a216fff] PGTABLE
BRK [0x44a217000, 0x44a217fff] PGTABLE
Secure boot disabled
RAMDISK: [mem 0x7f896000-0x7fffefff]
ACPI: Early table checksum verification disabled
ACPI: RSDP 0x00000000ACDFE014 000024 (v02 SuperCRA)
ACPI: XSDT 0x00000000ACDFE1C0 0000EC (v01 SuperCRA TP-N11   00001140 PTEC 00000002)
ACPI: FACP 0x00000000ACDF9000 00010C (v05 SuperCRA TP-N11   00001140 PTEC 00000002)
ACPI: DSDT 0x00000000ACDE3000 011204 (v02 SuperCRA TP-N11   00001140 INTL 20120711)
ACPI: FACS 0x00000000ACD68000 000040
ACPI: ASF! 0x00000000ACDFD000 0000A5 (v32 SuperCRA TP-N11   00001140 PTEC 00000002)
ACPI: HPET 0x00000000ACDFC000 000038 (v01 SuperCRA TP-N11   00001140 PTEC 00000002)
ACPI: ECDT 0x00000000ACDFB000 000052 (v01 SuperCRA TP-N11   00001140 PTEC 00000002)
ACPI: APIC 0x00000000ACDF8000 000098 (v01 SuperCRA TP-N11   00001140 PTEC 00000002)
ACPI: MCFG 0x00000000ACDF7000 00003C (v01 SuperCRA TP-N11   00001140 PTEC 00000002)
ACPI: SSDT 0x00000000ACDF6000 000033 (v01 SuperCRA TP-SSDT1 00000100 INTL 20120711)
ACPI: SSDT 0x00000000ACDF5000 000486 (v01 SuperCRA TP-SSDT2 00000200 INTL 20120711)
ACPI: SSDT 0x00000000ACDE2000 0009CB (v01 SuperCRA SataAhci 00001000 INTL 20120711)
ACPI: SSDT 0x00000000ACDE1000 000152 (v01 SuperCRA Rmv_Batt 00001000 INTL 20120711)
ACPI: SSDT 0x00000000ACDE0000 0006C5 (v01 SuperCRA Cpu0Ist  00003000 INTL 20120711)
ACPI: SSDT 0x00000000ACDDF000 000B74 (v02 SuperCRA CpuSsdt  00003000 INTL 20120711)
ACPI: SSDT 0x00000000ACDDE000 000369 (v02 SuperCRA CtdpB    00001000 INTL 20120711)
ACPI: SSDT 0x00000000ACDDC000 001475 (v01 SuperCRA SaSsdt   00003000 INTL 20120711)
ACPI: SSDT 0x00000000ACDDB000 000394 (v02 SuperCRA CppcTabl 00001000 INTL 20120711)
ACPI: PCCT 0x00000000ACDDA000 00006E (v05 SuperCRA TP-N11   00001140 PTEC 00000002)
ACPI: SSDT 0x00000000ACDD9000 000AC4 (v02 SuperCRA Cpc_Tabl 00001000 INTL 20120711)
ACPI: TCPA 0x00000000ACDD8000 000032 (v02 PTL    SuperCRA   06040000 LNVO 00000001)
ACPI: SSDT 0x00000000ACDD7000 0006B0 (v02 Intel_ TpmTable 00001000 INTL 20120711)
ACPI: UEFI 0x00000000ACDD6000 000042 (v01 SuperCRA TP-N11   00001140 PTEC 00000002)
ACPI: POAT 0x00000000ACCB0000 000055 (v03 SuperCRA TP-N11   00001140 PTEC 00000002)
ACPI: BATB 0x00000000ACDD5000 000046 (v01 SuperCRA TP-N11   00001140 PTEC 00000002)
ACPI: FPDT 0x00000000ACDD4000 000064 (v01 SuperCRA TP-N11   00001140 PTEC 00000002)
ACPI: UEFI 0x00000000ACDD3000 0002F6 (v01 SuperCRA TP-N11   00001140 PTEC 00000002)
ACPI: DMAR 0x00000000ACDD2000 0000B0 (v01 SuperCRA TP-N11   00001140 PTEC 00000002)
ACPI: Local APIC address 0xfee00000
No NUMA configuration found
Faking a node at [mem 0x0000000000000000-0x000000044dffffff]
NODE_DATA(0) allocated [mem 0x44dffc000-0x44dffffff]
Zone ranges:
  DMA      [mem 0x0000000000001000-0x0000000000ffffff]
  DMA32    [mem 0x0000000001000000-0x00000000ffffffff]
  Normal   [mem 0x0000000100000000-0x000000044dffffff]
  Device   empty
Movable zone start for each node
Early memory node ranges
  node   0: [mem 0x0000000000001000-0x0000000000057fff]
  node   0: [mem 0x0000000000059000-0x000000000008bfff]
  node   0: [mem 0x0000000000100000-0x00000000aa248fff]
  node   0: [mem 0x00000000acdff000-0x00000000acdfffff]
  node   0: [mem 0x0000000100000000-0x000000044dffffff]
Initmem setup node 0 [mem 0x0000000000001000-0x000000044dffffff]
On node 0 totalpages: 4162004
  DMA zone: 64 pages used for memmap
  DMA zone: 72 pages reserved
  DMA zone: 3978 pages, LIFO batch:0
  DMA32 zone: 10826 pages used for memmap
  DMA32 zone: 692810 pages, LIFO batch:31
  Normal zone: 54144 pages used for memmap
  Normal zone: 3465216 pages, LIFO batch:31
Reserving Intel graphics memory at 0x00000000ae000000-0x00000000afffffff
ACPI: PM-Timer IO Port: 0x1808
ACPI: Local APIC address 0xfee00000
ACPI: LAPIC_NMI (acpi_id[0x00] high edge lint[0x1])
ACPI: LAPIC_NMI (acpi_id[0x01] high edge lint[0x1])
IOAPIC[0]: apic_id 2, version 32, address 0xfec00000, GSI 0-39
ACPI: INT_SRC_OVR (bus 0 bus_irq 0 global_irq 2 dfl dfl)
ACPI: INT_SRC_OVR (bus 0 bus_irq 9 global_irq 9 high level)
ACPI: IRQ0 used by override.
ACPI: IRQ9 used by override.
Using ACPI (MADT) for SMP configuration information
ACPI: HPET id: 0x8086a301 base: 0xfed00000
smpboot: Allowing 8 CPUs, 4 hotplug CPUs
e820: [mem 0xb0000000-0xf80f7fff] available for PCI devices
Booting paravirtualized kernel on bare hardware
clocksource: refined-jiffies: mask: 0xffffffff max_cycles: 0xffffffff, max_idle_ns: 6370452778343963 ns
setup_percpu: NR_CPUS:128 nr_cpumask_bits:128 nr_cpu_ids:8 nr_node_ids:1
percpu: Embedded 36 pages/cpu @ffff979acdc00000 s107032 r8192 d32232 u262144
pcpu-alloc: s107032 r8192 d32232 u262144 alloc=1*2097152
pcpu-alloc: [0] 0 1 2 3 4 5 6 7
Built 1 zonelists, mobility grouping on.  Total pages: 4096898
Policy zone: Normal
PID hash table entries: 4096 (order: 3, 32768 bytes)
Calgary: detecting Calgary via BIOS EBDA area
Calgary: Unable to locate Rio Grande table in EBDA - bailing!
Memory: 16215128K/16648016K available (6641K kernel code, 1146K rwdata, 2784K rodata, 1312K init, 1008K bss, 432888K reserved, 0K cma-reserved)
SLUB: HWalign=64, Order=0-3, MinObjects=0, CPUs=8, Nodes=1
ftrace: allocating 29043 entries in 114 pages
Preemptible hierarchical RCU implementation.
        RCU restricting CPUs from NR_CPUS=128 to nr_cpu_ids=8.
        Tasks RCU enabled.
RCU: Adjusting geometry for rcu_fanout_leaf=16, nr_cpu_ids=8
NR_IRQS: 8448, nr_irqs: 760, preallocated irqs: 16
Console: colour dummy device 80x25
console [tty0] enabled
clocksource: hpet: mask: 0xffffffff max_cycles: 0xffffffff, max_idle_ns: 133484882848 ns
hpet clockevent registered
tsc: Fast TSC calibration using PIT
tsc: Detected 2194.919 MHz processor
Calibrating delay loop (skipped), value calculated using timer frequency.. 4391.74 BogoMIPS (lpj=7316396)
pid_max: default: 32768 minimum: 301
ACPI: Core revision 20170728
ACPI: 12 ACPI AML tables successfully acquired and loaded
Security Framework initialized
Yama: becoming mindful.
Dentry cache hash table entries: 2097152 (order: 12, 16777216 bytes)
Inode-cache hash table entries: 1048576 (order: 11, 8388608 bytes)
Mount-cache hash table entries: 32768 (order: 6, 262144 bytes)
Mountpoint-cache hash table entries: 32768 (order: 6, 262144 bytes)
CPU: Physical Processor ID: 0
CPU: Processor Core ID: 0
ENERGY_PERF_BIAS: Set to 'normal', was 'performance'
ENERGY_PERF_BIAS: View and update with x86_energy_perf_policy(8)
mce: CPU supports 7 MCE banks
CPU0: Thermal monitoring enabled (TM1)
process: using mwait in idle threads
Last level iTLB entries: 4KB 64, 2MB 8, 4MB 8
Last level dTLB entries: 4KB 64, 2MB 0, 4MB 0, 1GB 4
Freeing SMP alternatives memory: 24K
smpboot: Max logical packages: 4
DMAR: Host address width 39
DMAR: DRHD base: 0x000000fed90000 flags: 0x0
DMAR: dmar0: reg_base_addr fed90000 ver 1:0 cap 1c0000c40660462 ecap 7e1ff0505e
DMAR: DRHD base: 0x000000fed91000 flags: 0x1
DMAR: dmar1: reg_base_addr fed91000 ver 1:0 cap d2008c20660462 ecap f010da
DMAR: RMRR base: 0x000000aaf45000 end: 0x000000aaf5bfff
DMAR: RMRR base: 0x000000ad800000 end: 0x000000afffffff
DMAR-IR: IOAPIC id 2 under DRHD base  0xfed91000 IOMMU 1
DMAR-IR: HPET id 0 under DRHD base 0xfed91000
DMAR-IR: x2apic is disabled because BIOS sets x2apic opt out bit.
DMAR-IR: Use 'intremap=no_x2apic_optout' to override the BIOS setting.
DMAR-IR: Enabled IRQ remapping in xapic mode
x2apic: IRQ remapping doesn't support X2APIC mode
..TIMER: vector=0x30 apic1=0 pin1=2 apic2=-1 pin2=-1
TSC deadline timer enabled
Performance Events: PEBS fmt2+, Broadwell events, 16-deep LBR, full-width counters, Intel PMU driver.
... version:                3
... bit width:              48
... generic registers:      4
... value mask:             0000ffffffffffff
... max period:             00007fffffffffff
... fixed-purpose events:   3
... event mask:             000000070000000f
Hierarchical SRCU implementation.
NMI watchdog: Enabled. Permanently consumes one hw-PMU counter.
smp: Bringing up secondary CPUs ...
x86: Booting SMP configuration:
.... node  #0, CPUs:      #1 #2 #3
smp: Brought up 1 node, 4 CPUs
`

var BootLog = readLines(bootLog)
