from scapy.all import *
import sys
import os
import time



"""
main function of the script.
"""
def main():
    try:
        interface = input("[*] Enter desired interface: ")
        victimIP = input("[*] Enter victim IP: ")
        gateIP = input("[*] Enter Router IP: ")
    except KeyboardInterrupt:
        print("\n[*] Request shutdown")
        print("[*] Exiting...")
        sys.exit(1)
    
    print("\n[*] Enabling IP Forwarding...\n")
    os.system("echo 1 > /proc/sys/net/ipv4/ip_forward")


if __name__ == "__main__":
    main()