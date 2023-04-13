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


if __name__ == "__main__":
    main()