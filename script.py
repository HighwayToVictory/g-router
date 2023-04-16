from scapy.all import Ether, ARP, srp, send
import argparse
import time
import os
import sys



def _enable_linux_iproute():
    """
    Enables IP route ( IP Forward ) in linux-based distros.
    """
    file_path = "/proc/sys/net/ipv4/ip_forward"
    
    with open(file_path) as f:
        if f.read() == 1:
            return  # already enabled
    with open(file_path, "w") as f:
        print(1, file=f)
