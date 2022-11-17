// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.equinixmetal.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;


public final class VolumeSnapshotPolicyArgs extends com.pulumi.resources.ResourceArgs {

    public static final VolumeSnapshotPolicyArgs Empty = new VolumeSnapshotPolicyArgs();

    @Import(name="snapshotCount", required=true)
    private Output<Integer> snapshotCount;

    public Output<Integer> snapshotCount() {
        return this.snapshotCount;
    }

    @Import(name="snapshotFrequency", required=true)
    private Output<String> snapshotFrequency;

    public Output<String> snapshotFrequency() {
        return this.snapshotFrequency;
    }

    private VolumeSnapshotPolicyArgs() {}

    private VolumeSnapshotPolicyArgs(VolumeSnapshotPolicyArgs $) {
        this.snapshotCount = $.snapshotCount;
        this.snapshotFrequency = $.snapshotFrequency;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VolumeSnapshotPolicyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VolumeSnapshotPolicyArgs $;

        public Builder() {
            $ = new VolumeSnapshotPolicyArgs();
        }

        public Builder(VolumeSnapshotPolicyArgs defaults) {
            $ = new VolumeSnapshotPolicyArgs(Objects.requireNonNull(defaults));
        }

        public Builder snapshotCount(Output<Integer> snapshotCount) {
            $.snapshotCount = snapshotCount;
            return this;
        }

        public Builder snapshotCount(Integer snapshotCount) {
            return snapshotCount(Output.of(snapshotCount));
        }

        public Builder snapshotFrequency(Output<String> snapshotFrequency) {
            $.snapshotFrequency = snapshotFrequency;
            return this;
        }

        public Builder snapshotFrequency(String snapshotFrequency) {
            return snapshotFrequency(Output.of(snapshotFrequency));
        }

        public VolumeSnapshotPolicyArgs build() {
            $.snapshotCount = Objects.requireNonNull($.snapshotCount, "expected parameter 'snapshotCount' to be non-null");
            $.snapshotFrequency = Objects.requireNonNull($.snapshotFrequency, "expected parameter 'snapshotFrequency' to be non-null");
            return $;
        }
    }

}
