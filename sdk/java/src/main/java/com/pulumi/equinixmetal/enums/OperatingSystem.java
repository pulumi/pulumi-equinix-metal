// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.equinixmetal.enums;

import com.pulumi.core.annotations.EnumType;
import java.lang.String;
import java.util.Objects;
import java.util.StringJoiner;

    @EnumType
    public enum OperatingSystem {
        Alpine3("alpine_3"),
        CentOS6("centos_6"),
        CentOS7("centos_7"),
        CentOS8("centos_8"),
        CoreOSAlpha("coreos_alpha"),
        CoreOSBeta("coreos_beta"),
        CoreOSStable("coreos_stable"),
        CustomIPXE("custom_ipxe"),
        Debian8("debian_8"),
        Debian9("debian_9"),
        Debian10("debian_10"),
        FlatcarAlpha("flatcar_alpha"),
        FlatcarBeta("flatcar_beta"),
        FlatcarEdge("flatcar_edge"),
        FlatcarStable("flatcar_stable"),
        FreeBSD10_4("freebsd_10_4"),
        FreeBSD11_1("freebsd_11_1"),
        FreeBSD11_2("freebsd_11_2"),
        FreeBSD12Testing("freebsd_12_testing"),
        NixOS18_03("nixos_18_03"),
        NixOS19_03("nixos_19_03"),
        OpenSUSE42_3("opensuse_42_3"),
        RancherOS("rancher"),
        RHEL7("rhel_7"),
        RHEL8("rhel_8"),
        ScientificLinux6("scientific_6"),
        SLES12SP3("suse_sles12_sp3"),
        Ubuntu1404("ubuntu_14_04"),
        Ubuntu1604("ubuntu_16_04"),
        Ubuntu1710("ubuntu_17_10"),
        Ubuntu1804("ubuntu_18_04"),
        Ubuntu2004("ubuntu_20_04"),
        Ubuntu2010("ubuntu_20_10"),
        VMWareEsxi6_5("vmware_esxi_6_5"),
        VMWareEsxi6_7("vmware_esxi_6_7"),
        VMWareEsxi7_0("vmware_esxi_7_0"),
        Windows2012R2("windows_2012_r2"),
        Windows2016("windows_2016"),
        Windows2019("windows_2019");

        private final String value;

        OperatingSystem(String value) {
            this.value = Objects.requireNonNull(value);
        }

        @EnumType.Converter
        public String getValue() {
            return this.value;
        }

        @Override
        public String toString() {
            return new StringJoiner(", ", "OperatingSystem[", "]")
                .add("value='" + this.value + "'")
                .toString();
        }
    }
