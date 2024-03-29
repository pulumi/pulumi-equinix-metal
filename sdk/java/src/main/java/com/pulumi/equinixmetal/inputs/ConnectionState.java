// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.equinixmetal.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.equinixmetal.inputs.ConnectionPortArgs;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ConnectionState extends com.pulumi.resources.ResourceArgs {

    public static final ConnectionState Empty = new ConnectionState();

    /**
     * Description for the connection resource
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Description for the connection resource
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Facility where the connection will be created
     * 
     */
    @Import(name="facility")
    private @Nullable Output<String> facility;

    /**
     * @return Facility where the connection will be created
     * 
     */
    public Optional<Output<String>> facility() {
        return Optional.ofNullable(this.facility);
    }

    /**
     * Metro where the connection will be created
     * 
     */
    @Import(name="metro")
    private @Nullable Output<String> metro;

    /**
     * @return Metro where the connection will be created
     * 
     */
    public Optional<Output<String>> metro() {
        return Optional.ofNullable(this.metro);
    }

    /**
     * Mode for connections in IBX facilities with the dedicated type - standard or tunnel
     * 
     */
    @Import(name="mode")
    private @Nullable Output<String> mode;

    /**
     * @return Mode for connections in IBX facilities with the dedicated type - standard or tunnel
     * 
     */
    public Optional<Output<String>> mode() {
        return Optional.ofNullable(this.mode);
    }

    /**
     * Name of the connection resource
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the connection resource
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * ID of the organization responsible for the connection
     * 
     */
    @Import(name="organizationId")
    private @Nullable Output<String> organizationId;

    /**
     * @return ID of the organization responsible for the connection
     * 
     */
    public Optional<Output<String>> organizationId() {
        return Optional.ofNullable(this.organizationId);
    }

    /**
     * List of connection ports - primary (`ports[0]`) and secondary (`ports[1]`). Schema of port is described in documentation of the equinix-metal.Connection datasource.
     * 
     */
    @Import(name="ports")
    private @Nullable Output<List<ConnectionPortArgs>> ports;

    /**
     * @return List of connection ports - primary (`ports[0]`) and secondary (`ports[1]`). Schema of port is described in documentation of the equinix-metal.Connection datasource.
     * 
     */
    public Optional<Output<List<ConnectionPortArgs>>> ports() {
        return Optional.ofNullable(this.ports);
    }

    /**
     * ID of the project where the connection is scoped to, must be set for shared connection
     * 
     */
    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    /**
     * @return ID of the project where the connection is scoped to, must be set for shared connection
     * 
     */
    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    /**
     * Connection redundancy - redundant or primary
     * 
     */
    @Import(name="redundancy")
    private @Nullable Output<String> redundancy;

    /**
     * @return Connection redundancy - redundant or primary
     * 
     */
    public Optional<Output<String>> redundancy() {
        return Optional.ofNullable(this.redundancy);
    }

    /**
     * Port speed in bits per second
     * 
     */
    @Import(name="speed")
    private @Nullable Output<Integer> speed;

    /**
     * @return Port speed in bits per second
     * 
     */
    public Optional<Output<Integer>> speed() {
        return Optional.ofNullable(this.speed);
    }

    /**
     * Status of the connection resource
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return Status of the connection resource
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * String list of tags
     * 
     */
    @Import(name="tags")
    private @Nullable Output<List<String>> tags;

    /**
     * @return String list of tags
     * 
     */
    public Optional<Output<List<String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * Fabric Token from the [Equinix Fabric Portal](https://ecxfabric.equinix.com/dashboard)
     * 
     */
    @Import(name="token")
    private @Nullable Output<String> token;

    /**
     * @return Fabric Token from the [Equinix Fabric Portal](https://ecxfabric.equinix.com/dashboard)
     * 
     */
    public Optional<Output<String>> token() {
        return Optional.ofNullable(this.token);
    }

    /**
     * Connection type - dedicated or shared
     * 
     */
    @Import(name="type")
    private @Nullable Output<String> type;

    /**
     * @return Connection type - dedicated or shared
     * 
     */
    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    private ConnectionState() {}

    private ConnectionState(ConnectionState $) {
        this.description = $.description;
        this.facility = $.facility;
        this.metro = $.metro;
        this.mode = $.mode;
        this.name = $.name;
        this.organizationId = $.organizationId;
        this.ports = $.ports;
        this.projectId = $.projectId;
        this.redundancy = $.redundancy;
        this.speed = $.speed;
        this.status = $.status;
        this.tags = $.tags;
        this.token = $.token;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ConnectionState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ConnectionState $;

        public Builder() {
            $ = new ConnectionState();
        }

        public Builder(ConnectionState defaults) {
            $ = new ConnectionState(Objects.requireNonNull(defaults));
        }

        /**
         * @param description Description for the connection resource
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Description for the connection resource
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param facility Facility where the connection will be created
         * 
         * @return builder
         * 
         */
        public Builder facility(@Nullable Output<String> facility) {
            $.facility = facility;
            return this;
        }

        /**
         * @param facility Facility where the connection will be created
         * 
         * @return builder
         * 
         */
        public Builder facility(String facility) {
            return facility(Output.of(facility));
        }

        /**
         * @param metro Metro where the connection will be created
         * 
         * @return builder
         * 
         */
        public Builder metro(@Nullable Output<String> metro) {
            $.metro = metro;
            return this;
        }

        /**
         * @param metro Metro where the connection will be created
         * 
         * @return builder
         * 
         */
        public Builder metro(String metro) {
            return metro(Output.of(metro));
        }

        /**
         * @param mode Mode for connections in IBX facilities with the dedicated type - standard or tunnel
         * 
         * @return builder
         * 
         */
        public Builder mode(@Nullable Output<String> mode) {
            $.mode = mode;
            return this;
        }

        /**
         * @param mode Mode for connections in IBX facilities with the dedicated type - standard or tunnel
         * 
         * @return builder
         * 
         */
        public Builder mode(String mode) {
            return mode(Output.of(mode));
        }

        /**
         * @param name Name of the connection resource
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the connection resource
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param organizationId ID of the organization responsible for the connection
         * 
         * @return builder
         * 
         */
        public Builder organizationId(@Nullable Output<String> organizationId) {
            $.organizationId = organizationId;
            return this;
        }

        /**
         * @param organizationId ID of the organization responsible for the connection
         * 
         * @return builder
         * 
         */
        public Builder organizationId(String organizationId) {
            return organizationId(Output.of(organizationId));
        }

        /**
         * @param ports List of connection ports - primary (`ports[0]`) and secondary (`ports[1]`). Schema of port is described in documentation of the equinix-metal.Connection datasource.
         * 
         * @return builder
         * 
         */
        public Builder ports(@Nullable Output<List<ConnectionPortArgs>> ports) {
            $.ports = ports;
            return this;
        }

        /**
         * @param ports List of connection ports - primary (`ports[0]`) and secondary (`ports[1]`). Schema of port is described in documentation of the equinix-metal.Connection datasource.
         * 
         * @return builder
         * 
         */
        public Builder ports(List<ConnectionPortArgs> ports) {
            return ports(Output.of(ports));
        }

        /**
         * @param ports List of connection ports - primary (`ports[0]`) and secondary (`ports[1]`). Schema of port is described in documentation of the equinix-metal.Connection datasource.
         * 
         * @return builder
         * 
         */
        public Builder ports(ConnectionPortArgs... ports) {
            return ports(List.of(ports));
        }

        /**
         * @param projectId ID of the project where the connection is scoped to, must be set for shared connection
         * 
         * @return builder
         * 
         */
        public Builder projectId(@Nullable Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId ID of the project where the connection is scoped to, must be set for shared connection
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param redundancy Connection redundancy - redundant or primary
         * 
         * @return builder
         * 
         */
        public Builder redundancy(@Nullable Output<String> redundancy) {
            $.redundancy = redundancy;
            return this;
        }

        /**
         * @param redundancy Connection redundancy - redundant or primary
         * 
         * @return builder
         * 
         */
        public Builder redundancy(String redundancy) {
            return redundancy(Output.of(redundancy));
        }

        /**
         * @param speed Port speed in bits per second
         * 
         * @return builder
         * 
         */
        public Builder speed(@Nullable Output<Integer> speed) {
            $.speed = speed;
            return this;
        }

        /**
         * @param speed Port speed in bits per second
         * 
         * @return builder
         * 
         */
        public Builder speed(Integer speed) {
            return speed(Output.of(speed));
        }

        /**
         * @param status Status of the connection resource
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status Status of the connection resource
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param tags String list of tags
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<List<String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags String list of tags
         * 
         * @return builder
         * 
         */
        public Builder tags(List<String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tags String list of tags
         * 
         * @return builder
         * 
         */
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }

        /**
         * @param token Fabric Token from the [Equinix Fabric Portal](https://ecxfabric.equinix.com/dashboard)
         * 
         * @return builder
         * 
         */
        public Builder token(@Nullable Output<String> token) {
            $.token = token;
            return this;
        }

        /**
         * @param token Fabric Token from the [Equinix Fabric Portal](https://ecxfabric.equinix.com/dashboard)
         * 
         * @return builder
         * 
         */
        public Builder token(String token) {
            return token(Output.of(token));
        }

        /**
         * @param type Connection type - dedicated or shared
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type Connection type - dedicated or shared
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public ConnectionState build() {
            return $;
        }
    }

}
