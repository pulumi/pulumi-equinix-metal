// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.equinixmetal.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ConnectionPort {
    private @Nullable String id;
    private @Nullable String linkStatus;
    /**
     * @return Name of the connection resource
     * 
     */
    private @Nullable String name;
    private @Nullable String role;
    /**
     * @return Port speed in bits per second
     * 
     */
    private @Nullable Integer speed;
    /**
     * @return Status of the connection resource
     * 
     */
    private @Nullable String status;
    private @Nullable List<Object> virtualCircuitIds;

    private ConnectionPort() {}
    public Optional<String> id() {
        return Optional.ofNullable(this.id);
    }
    public Optional<String> linkStatus() {
        return Optional.ofNullable(this.linkStatus);
    }
    /**
     * @return Name of the connection resource
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }
    public Optional<String> role() {
        return Optional.ofNullable(this.role);
    }
    /**
     * @return Port speed in bits per second
     * 
     */
    public Optional<Integer> speed() {
        return Optional.ofNullable(this.speed);
    }
    /**
     * @return Status of the connection resource
     * 
     */
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }
    public List<Object> virtualCircuitIds() {
        return this.virtualCircuitIds == null ? List.of() : this.virtualCircuitIds;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ConnectionPort defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String id;
        private @Nullable String linkStatus;
        private @Nullable String name;
        private @Nullable String role;
        private @Nullable Integer speed;
        private @Nullable String status;
        private @Nullable List<Object> virtualCircuitIds;
        public Builder() {}
        public Builder(ConnectionPort defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.linkStatus = defaults.linkStatus;
    	      this.name = defaults.name;
    	      this.role = defaults.role;
    	      this.speed = defaults.speed;
    	      this.status = defaults.status;
    	      this.virtualCircuitIds = defaults.virtualCircuitIds;
        }

        @CustomType.Setter
        public Builder id(@Nullable String id) {
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder linkStatus(@Nullable String linkStatus) {
            this.linkStatus = linkStatus;
            return this;
        }
        @CustomType.Setter
        public Builder name(@Nullable String name) {
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder role(@Nullable String role) {
            this.role = role;
            return this;
        }
        @CustomType.Setter
        public Builder speed(@Nullable Integer speed) {
            this.speed = speed;
            return this;
        }
        @CustomType.Setter
        public Builder status(@Nullable String status) {
            this.status = status;
            return this;
        }
        @CustomType.Setter
        public Builder virtualCircuitIds(@Nullable List<Object> virtualCircuitIds) {
            this.virtualCircuitIds = virtualCircuitIds;
            return this;
        }
        public Builder virtualCircuitIds(Object... virtualCircuitIds) {
            return virtualCircuitIds(List.of(virtualCircuitIds));
        }
        public ConnectionPort build() {
            final var o = new ConnectionPort();
            o.id = id;
            o.linkStatus = linkStatus;
            o.name = name;
            o.role = role;
            o.speed = speed;
            o.status = status;
            o.virtualCircuitIds = virtualCircuitIds;
            return o;
        }
    }
}
