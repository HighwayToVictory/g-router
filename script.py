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


def _enable_windows_iproute():
    """
    Enables IP route (IP Forwarding) in Windows.
    """
    from windows.service import WService
    # enable Remote Access service
    service = WService("RemoteAccess")
    service.start()


def enable_ip_route(verbose=True):
    """
    Enables IP forwarding.
    """
    if verbose:
        print("[!] Enabling IP Routing...")
    _enable_windows_iproute() if "nt" in os.name else _enable_linux_iproute()
    if verbose:
        print("[!] IP Routing enabled.")