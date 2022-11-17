// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.equinixmetal.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetVirtualCircuitResult {
    /**
     * @return Description for the Virtual Circuit resource
     * 
     */
    private String description;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return Name of the virtual circuit resource
     * 
     */
    private String name;
    private Integer nniVlan;
    private Integer nniVnid;
    /**
     * @return ID of project to which the VC belongs
     * * `vnid`, `nni_vlan`, `nni_nvid` - VLAN parameters, see the [documentation for Equinix Fabric](https://metal.equinix.com/developers/docs/networking/fabric/)
     * 
     */
    private String projectId;
    /**
     * @return Status of the virtal circuit
     * 
     */
    private String status;
    /**
     * @return Tags for the Virtual Circuit resource
     * 
     */
    private List<String> tags;
    private String virtualCircuitId;
    private Integer vnid;

    private GetVirtualCircuitResult() {}
    /**
     * @return Description for the Virtual Circuit resource
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Name of the virtual circuit resource
     * 
     */
    public String name() {
        return this.name;
    }
    public Integer nniVlan() {
        return this.nniVlan;
    }
    public Integer nniVnid() {
        return this.nniVnid;
    }
    /**
     * @return ID of project to which the VC belongs
     * * `vnid`, `nni_vlan`, `nni_nvid` - VLAN parameters, see the [documentation for Equinix Fabric](https://metal.equinix.com/developers/docs/networking/fabric/)
     * 
     */
    public String projectId() {
        return this.projectId;
    }
    /**
     * @return Status of the virtal circuit
     * 
     */
    public String status() {
        return this.status;
    }
    /**
     * @return Tags for the Virtual Circuit resource
     * 
     */
    public List<String> tags() {
        return this.tags;
    }
    public String virtualCircuitId() {
        return this.virtualCircuitId;
    }
    public Integer vnid() {
        return this.vnid;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetVirtualCircuitResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String description;
        private String id;
        private String name;
        private Integer nniVlan;
        private Integer nniVnid;
        private String projectId;
        private String status;
        private List<String> tags;
        private String virtualCircuitId;
        private Integer vnid;
        public Builder() {}
        public Builder(GetVirtualCircuitResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.description = defaults.description;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.nniVlan = defaults.nniVlan;
    	      this.nniVnid = defaults.nniVnid;
    	      this.projectId = defaults.projectId;
    	      this.status = defaults.status;
    	      this.tags = defaults.tags;
    	      this.virtualCircuitId = defaults.virtualCircuitId;
    	      this.vnid = defaults.vnid;
        }

        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder nniVlan(Integer nniVlan) {
            this.nniVlan = Objects.requireNonNull(nniVlan);
            return this;
        }
        @CustomType.Setter
        public Builder nniVnid(Integer nniVnid) {
            this.nniVnid = Objects.requireNonNull(nniVnid);
            return this;
        }
        @CustomType.Setter
        public Builder projectId(String projectId) {
            this.projectId = Objects.requireNonNull(projectId);
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            this.status = Objects.requireNonNull(status);
            return this;
        }
        @CustomType.Setter
        public Builder tags(List<String> tags) {
            this.tags = Objects.requireNonNull(tags);
            return this;
        }
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }
        @CustomType.Setter
        public Builder virtualCircuitId(String virtualCircuitId) {
            this.virtualCircuitId = Objects.requireNonNull(virtualCircuitId);
            return this;
        }
        @CustomType.Setter
        public Builder vnid(Integer vnid) {
            this.vnid = Objects.requireNonNull(vnid);
            return this;
        }
        public GetVirtualCircuitResult build() {
            final var o = new GetVirtualCircuitResult();
            o.description = description;
            o.id = id;
            o.name = name;
            o.nniVlan = nniVlan;
            o.nniVnid = nniVnid;
            o.projectId = projectId;
            o.status = status;
            o.tags = tags;
            o.virtualCircuitId = virtualCircuitId;
            o.vnid = vnid;
            return o;
        }
    }
}
