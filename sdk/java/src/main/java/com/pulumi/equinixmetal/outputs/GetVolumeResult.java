// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.equinixmetal.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.equinixmetal.outputs.GetVolumeSnapshotPolicy;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetVolumeResult {
    private String billingCycle;
    private String created;
    private String description;
    private List<String> deviceIds;
    private String facility;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private Boolean locked;
    private String name;
    private String plan;
    private String projectId;
    private Integer size;
    private List<GetVolumeSnapshotPolicy> snapshotPolicies;
    private String state;
    private String updated;
    private String volumeId;

    private GetVolumeResult() {}
    public String billingCycle() {
        return this.billingCycle;
    }
    public String created() {
        return this.created;
    }
    public String description() {
        return this.description;
    }
    public List<String> deviceIds() {
        return this.deviceIds;
    }
    public String facility() {
        return this.facility;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public Boolean locked() {
        return this.locked;
    }
    public String name() {
        return this.name;
    }
    public String plan() {
        return this.plan;
    }
    public String projectId() {
        return this.projectId;
    }
    public Integer size() {
        return this.size;
    }
    public List<GetVolumeSnapshotPolicy> snapshotPolicies() {
        return this.snapshotPolicies;
    }
    public String state() {
        return this.state;
    }
    public String updated() {
        return this.updated;
    }
    public String volumeId() {
        return this.volumeId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetVolumeResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String billingCycle;
        private String created;
        private String description;
        private List<String> deviceIds;
        private String facility;
        private String id;
        private Boolean locked;
        private String name;
        private String plan;
        private String projectId;
        private Integer size;
        private List<GetVolumeSnapshotPolicy> snapshotPolicies;
        private String state;
        private String updated;
        private String volumeId;
        public Builder() {}
        public Builder(GetVolumeResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.billingCycle = defaults.billingCycle;
    	      this.created = defaults.created;
    	      this.description = defaults.description;
    	      this.deviceIds = defaults.deviceIds;
    	      this.facility = defaults.facility;
    	      this.id = defaults.id;
    	      this.locked = defaults.locked;
    	      this.name = defaults.name;
    	      this.plan = defaults.plan;
    	      this.projectId = defaults.projectId;
    	      this.size = defaults.size;
    	      this.snapshotPolicies = defaults.snapshotPolicies;
    	      this.state = defaults.state;
    	      this.updated = defaults.updated;
    	      this.volumeId = defaults.volumeId;
        }

        @CustomType.Setter
        public Builder billingCycle(String billingCycle) {
            this.billingCycle = Objects.requireNonNull(billingCycle);
            return this;
        }
        @CustomType.Setter
        public Builder created(String created) {
            this.created = Objects.requireNonNull(created);
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder deviceIds(List<String> deviceIds) {
            this.deviceIds = Objects.requireNonNull(deviceIds);
            return this;
        }
        public Builder deviceIds(String... deviceIds) {
            return deviceIds(List.of(deviceIds));
        }
        @CustomType.Setter
        public Builder facility(String facility) {
            this.facility = Objects.requireNonNull(facility);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder locked(Boolean locked) {
            this.locked = Objects.requireNonNull(locked);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder plan(String plan) {
            this.plan = Objects.requireNonNull(plan);
            return this;
        }
        @CustomType.Setter
        public Builder projectId(String projectId) {
            this.projectId = Objects.requireNonNull(projectId);
            return this;
        }
        @CustomType.Setter
        public Builder size(Integer size) {
            this.size = Objects.requireNonNull(size);
            return this;
        }
        @CustomType.Setter
        public Builder snapshotPolicies(List<GetVolumeSnapshotPolicy> snapshotPolicies) {
            this.snapshotPolicies = Objects.requireNonNull(snapshotPolicies);
            return this;
        }
        public Builder snapshotPolicies(GetVolumeSnapshotPolicy... snapshotPolicies) {
            return snapshotPolicies(List.of(snapshotPolicies));
        }
        @CustomType.Setter
        public Builder state(String state) {
            this.state = Objects.requireNonNull(state);
            return this;
        }
        @CustomType.Setter
        public Builder updated(String updated) {
            this.updated = Objects.requireNonNull(updated);
            return this;
        }
        @CustomType.Setter
        public Builder volumeId(String volumeId) {
            this.volumeId = Objects.requireNonNull(volumeId);
            return this;
        }
        public GetVolumeResult build() {
            final var o = new GetVolumeResult();
            o.billingCycle = billingCycle;
            o.created = created;
            o.description = description;
            o.deviceIds = deviceIds;
            o.facility = facility;
            o.id = id;
            o.locked = locked;
            o.name = name;
            o.plan = plan;
            o.projectId = projectId;
            o.size = size;
            o.snapshotPolicies = snapshotPolicies;
            o.state = state;
            o.updated = updated;
            o.volumeId = volumeId;
            return o;
        }
    }
}