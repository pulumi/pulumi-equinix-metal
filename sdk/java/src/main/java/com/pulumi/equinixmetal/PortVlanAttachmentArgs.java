// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.equinixmetal;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PortVlanAttachmentArgs extends com.pulumi.resources.ResourceArgs {

    public static final PortVlanAttachmentArgs Empty = new PortVlanAttachmentArgs();

    /**
     * ID of device to be assigned to the VLAN
     * 
     */
    @Import(name="deviceId", required=true)
    private Output<String> deviceId;

    /**
     * @return ID of device to be assigned to the VLAN
     * 
     */
    public Output<String> deviceId() {
        return this.deviceId;
    }

    /**
     * Add port back to the bond when this resource is removed. Default is false.
     * 
     */
    @Import(name="forceBond")
    private @Nullable Output<Boolean> forceBond;

    /**
     * @return Add port back to the bond when this resource is removed. Default is false.
     * 
     */
    public Optional<Output<Boolean>> forceBond() {
        return Optional.ofNullable(this.forceBond);
    }

    /**
     * Mark this VLAN a native VLAN on the port. This can be used only if this assignment assigns second or further VLAN to the port. To ensure that this attachment is not first on a port, you can use `depends_on` pointing to another metal_port_vlan_attachment, just like in the layer2-individual example above.
     * 
     */
    @Import(name="native")
    private @Nullable Output<Boolean> native_;

    /**
     * @return Mark this VLAN a native VLAN on the port. This can be used only if this assignment assigns second or further VLAN to the port. To ensure that this attachment is not first on a port, you can use `depends_on` pointing to another metal_port_vlan_attachment, just like in the layer2-individual example above.
     * 
     */
    public Optional<Output<Boolean>> native_() {
        return Optional.ofNullable(this.native_);
    }

    /**
     * Name of network port to be assigned to the VLAN
     * 
     */
    @Import(name="portName", required=true)
    private Output<String> portName;

    /**
     * @return Name of network port to be assigned to the VLAN
     * 
     */
    public Output<String> portName() {
        return this.portName;
    }

    /**
     * VXLAN Network Identifier, integer
     * 
     */
    @Import(name="vlanVnid", required=true)
    private Output<Integer> vlanVnid;

    /**
     * @return VXLAN Network Identifier, integer
     * 
     */
    public Output<Integer> vlanVnid() {
        return this.vlanVnid;
    }

    private PortVlanAttachmentArgs() {}

    private PortVlanAttachmentArgs(PortVlanAttachmentArgs $) {
        this.deviceId = $.deviceId;
        this.forceBond = $.forceBond;
        this.native_ = $.native_;
        this.portName = $.portName;
        this.vlanVnid = $.vlanVnid;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PortVlanAttachmentArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PortVlanAttachmentArgs $;

        public Builder() {
            $ = new PortVlanAttachmentArgs();
        }

        public Builder(PortVlanAttachmentArgs defaults) {
            $ = new PortVlanAttachmentArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param deviceId ID of device to be assigned to the VLAN
         * 
         * @return builder
         * 
         */
        public Builder deviceId(Output<String> deviceId) {
            $.deviceId = deviceId;
            return this;
        }

        /**
         * @param deviceId ID of device to be assigned to the VLAN
         * 
         * @return builder
         * 
         */
        public Builder deviceId(String deviceId) {
            return deviceId(Output.of(deviceId));
        }

        /**
         * @param forceBond Add port back to the bond when this resource is removed. Default is false.
         * 
         * @return builder
         * 
         */
        public Builder forceBond(@Nullable Output<Boolean> forceBond) {
            $.forceBond = forceBond;
            return this;
        }

        /**
         * @param forceBond Add port back to the bond when this resource is removed. Default is false.
         * 
         * @return builder
         * 
         */
        public Builder forceBond(Boolean forceBond) {
            return forceBond(Output.of(forceBond));
        }

        /**
         * @param native_ Mark this VLAN a native VLAN on the port. This can be used only if this assignment assigns second or further VLAN to the port. To ensure that this attachment is not first on a port, you can use `depends_on` pointing to another metal_port_vlan_attachment, just like in the layer2-individual example above.
         * 
         * @return builder
         * 
         */
        public Builder native_(@Nullable Output<Boolean> native_) {
            $.native_ = native_;
            return this;
        }

        /**
         * @param native_ Mark this VLAN a native VLAN on the port. This can be used only if this assignment assigns second or further VLAN to the port. To ensure that this attachment is not first on a port, you can use `depends_on` pointing to another metal_port_vlan_attachment, just like in the layer2-individual example above.
         * 
         * @return builder
         * 
         */
        public Builder native_(Boolean native_) {
            return native_(Output.of(native_));
        }

        /**
         * @param portName Name of network port to be assigned to the VLAN
         * 
         * @return builder
         * 
         */
        public Builder portName(Output<String> portName) {
            $.portName = portName;
            return this;
        }

        /**
         * @param portName Name of network port to be assigned to the VLAN
         * 
         * @return builder
         * 
         */
        public Builder portName(String portName) {
            return portName(Output.of(portName));
        }

        /**
         * @param vlanVnid VXLAN Network Identifier, integer
         * 
         * @return builder
         * 
         */
        public Builder vlanVnid(Output<Integer> vlanVnid) {
            $.vlanVnid = vlanVnid;
            return this;
        }

        /**
         * @param vlanVnid VXLAN Network Identifier, integer
         * 
         * @return builder
         * 
         */
        public Builder vlanVnid(Integer vlanVnid) {
            return vlanVnid(Output.of(vlanVnid));
        }

        public PortVlanAttachmentArgs build() {
            $.deviceId = Objects.requireNonNull($.deviceId, "expected parameter 'deviceId' to be non-null");
            $.portName = Objects.requireNonNull($.portName, "expected parameter 'portName' to be non-null");
            $.vlanVnid = Objects.requireNonNull($.vlanVnid, "expected parameter 'vlanVnid' to be non-null");
            return $;
        }
    }

}
