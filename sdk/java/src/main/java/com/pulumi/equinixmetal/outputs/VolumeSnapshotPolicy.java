// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.equinixmetal.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class VolumeSnapshotPolicy {
    private Integer snapshotCount;
    private String snapshotFrequency;

    private VolumeSnapshotPolicy() {}
    public Integer snapshotCount() {
        return this.snapshotCount;
    }
    public String snapshotFrequency() {
        return this.snapshotFrequency;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(VolumeSnapshotPolicy defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Integer snapshotCount;
        private String snapshotFrequency;
        public Builder() {}
        public Builder(VolumeSnapshotPolicy defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.snapshotCount = defaults.snapshotCount;
    	      this.snapshotFrequency = defaults.snapshotFrequency;
        }

        @CustomType.Setter
        public Builder snapshotCount(Integer snapshotCount) {
            this.snapshotCount = Objects.requireNonNull(snapshotCount);
            return this;
        }
        @CustomType.Setter
        public Builder snapshotFrequency(String snapshotFrequency) {
            this.snapshotFrequency = Objects.requireNonNull(snapshotFrequency);
            return this;
        }
        public VolumeSnapshotPolicy build() {
            final var o = new VolumeSnapshotPolicy();
            o.snapshotCount = snapshotCount;
            o.snapshotFrequency = snapshotFrequency;
            return o;
        }
    }
}
