// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.equinixmetal.inputs;

import com.pulumi.core.Either;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.equinixmetal.enums.Facility;
import com.pulumi.equinixmetal.enums.IpBlockType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ReservedIpBlockState extends com.pulumi.resources.ResourceArgs {

    public static final ReservedIpBlockState Empty = new ReservedIpBlockState();

    @Import(name="address")
    private @Nullable Output<String> address;

    public Optional<Output<String>> address() {
        return Optional.ofNullable(this.address);
    }

    /**
     * Address family as integer (4 or 6)
     * 
     */
    @Import(name="addressFamily")
    private @Nullable Output<Integer> addressFamily;

    /**
     * @return Address family as integer (4 or 6)
     * 
     */
    public Optional<Output<Integer>> addressFamily() {
        return Optional.ofNullable(this.addressFamily);
    }

    /**
     * length of CIDR prefix of the block as integer
     * 
     */
    @Import(name="cidr")
    private @Nullable Output<Integer> cidr;

    /**
     * @return length of CIDR prefix of the block as integer
     * 
     */
    public Optional<Output<Integer>> cidr() {
        return Optional.ofNullable(this.cidr);
    }

    /**
     * Address and mask in CIDR notation, e.g. &#34;147.229.15.30/31&#34;
     * 
     */
    @Import(name="cidrNotation")
    private @Nullable Output<String> cidrNotation;

    /**
     * @return Address and mask in CIDR notation, e.g. &#34;147.229.15.30/31&#34;
     * 
     */
    public Optional<Output<String>> cidrNotation() {
        return Optional.ofNullable(this.cidrNotation);
    }

    /**
     * Arbitrary description
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Arbitrary description
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Facility where to allocate the public IP address block, makes sense only for type==public_ipv4, must be empty for type==global_ipv4, conflicts with `metro`
     * 
     */
    @Import(name="facility")
    private @Nullable Output<Either<String,Facility>> facility;

    /**
     * @return Facility where to allocate the public IP address block, makes sense only for type==public_ipv4, must be empty for type==global_ipv4, conflicts with `metro`
     * 
     */
    public Optional<Output<Either<String,Facility>>> facility() {
        return Optional.ofNullable(this.facility);
    }

    @Import(name="gateway")
    private @Nullable Output<String> gateway;

    public Optional<Output<String>> gateway() {
        return Optional.ofNullable(this.gateway);
    }

    /**
     * boolean flag whether addresses from a block are global (i.e. can be assigned in any facility)
     * 
     */
    @Import(name="global")
    private @Nullable Output<Boolean> global;

    /**
     * @return boolean flag whether addresses from a block are global (i.e. can be assigned in any facility)
     * 
     */
    public Optional<Output<Boolean>> global() {
        return Optional.ofNullable(this.global);
    }

    @Import(name="manageable")
    private @Nullable Output<Boolean> manageable;

    public Optional<Output<Boolean>> manageable() {
        return Optional.ofNullable(this.manageable);
    }

    @Import(name="management")
    private @Nullable Output<Boolean> management;

    public Optional<Output<Boolean>> management() {
        return Optional.ofNullable(this.management);
    }

    /**
     * Metro where to allocate the public IP address block, makes sense only for type==public_ipv4, must be empty for type==global_ipv4, conflicts with `facility`
     * 
     */
    @Import(name="metro")
    private @Nullable Output<String> metro;

    /**
     * @return Metro where to allocate the public IP address block, makes sense only for type==public_ipv4, must be empty for type==global_ipv4, conflicts with `facility`
     * 
     */
    public Optional<Output<String>> metro() {
        return Optional.ofNullable(this.metro);
    }

    /**
     * Mask in decimal notation, e.g. &#34;255.255.255.0&#34;
     * 
     */
    @Import(name="netmask")
    private @Nullable Output<String> netmask;

    /**
     * @return Mask in decimal notation, e.g. &#34;255.255.255.0&#34;
     * 
     */
    public Optional<Output<String>> netmask() {
        return Optional.ofNullable(this.netmask);
    }

    /**
     * Network IP address portion of the block specification
     * 
     */
    @Import(name="network")
    private @Nullable Output<String> network;

    /**
     * @return Network IP address portion of the block specification
     * 
     */
    public Optional<Output<String>> network() {
        return Optional.ofNullable(this.network);
    }

    /**
     * The metal project ID where to allocate the address block
     * 
     */
    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    /**
     * @return The metal project ID where to allocate the address block
     * 
     */
    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    /**
     * boolean flag whether addresses from a block are public
     * 
     */
    @Import(name="public")
    private @Nullable Output<Boolean> public_;

    /**
     * @return boolean flag whether addresses from a block are public
     * 
     */
    public Optional<Output<Boolean>> public_() {
        return Optional.ofNullable(this.public_);
    }

    /**
     * The number of allocated /32 addresses, a power of 2
     * 
     */
    @Import(name="quantity")
    private @Nullable Output<Integer> quantity;

    /**
     * @return The number of allocated /32 addresses, a power of 2
     * 
     */
    public Optional<Output<Integer>> quantity() {
        return Optional.ofNullable(this.quantity);
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
     * Either &#34;global_ipv4&#34; or &#34;public_ipv4&#34;, defaults to &#34;public_ipv4&#34; for backward compatibility
     * 
     */
    @Import(name="type")
    private @Nullable Output<Either<String,IpBlockType>> type;

    /**
     * @return Either &#34;global_ipv4&#34; or &#34;public_ipv4&#34;, defaults to &#34;public_ipv4&#34; for backward compatibility
     * 
     */
    public Optional<Output<Either<String,IpBlockType>>> type() {
        return Optional.ofNullable(this.type);
    }

    private ReservedIpBlockState() {}

    private ReservedIpBlockState(ReservedIpBlockState $) {
        this.address = $.address;
        this.addressFamily = $.addressFamily;
        this.cidr = $.cidr;
        this.cidrNotation = $.cidrNotation;
        this.description = $.description;
        this.facility = $.facility;
        this.gateway = $.gateway;
        this.global = $.global;
        this.manageable = $.manageable;
        this.management = $.management;
        this.metro = $.metro;
        this.netmask = $.netmask;
        this.network = $.network;
        this.projectId = $.projectId;
        this.public_ = $.public_;
        this.quantity = $.quantity;
        this.tags = $.tags;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ReservedIpBlockState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ReservedIpBlockState $;

        public Builder() {
            $ = new ReservedIpBlockState();
        }

        public Builder(ReservedIpBlockState defaults) {
            $ = new ReservedIpBlockState(Objects.requireNonNull(defaults));
        }

        public Builder address(@Nullable Output<String> address) {
            $.address = address;
            return this;
        }

        public Builder address(String address) {
            return address(Output.of(address));
        }

        /**
         * @param addressFamily Address family as integer (4 or 6)
         * 
         * @return builder
         * 
         */
        public Builder addressFamily(@Nullable Output<Integer> addressFamily) {
            $.addressFamily = addressFamily;
            return this;
        }

        /**
         * @param addressFamily Address family as integer (4 or 6)
         * 
         * @return builder
         * 
         */
        public Builder addressFamily(Integer addressFamily) {
            return addressFamily(Output.of(addressFamily));
        }

        /**
         * @param cidr length of CIDR prefix of the block as integer
         * 
         * @return builder
         * 
         */
        public Builder cidr(@Nullable Output<Integer> cidr) {
            $.cidr = cidr;
            return this;
        }

        /**
         * @param cidr length of CIDR prefix of the block as integer
         * 
         * @return builder
         * 
         */
        public Builder cidr(Integer cidr) {
            return cidr(Output.of(cidr));
        }

        /**
         * @param cidrNotation Address and mask in CIDR notation, e.g. &#34;147.229.15.30/31&#34;
         * 
         * @return builder
         * 
         */
        public Builder cidrNotation(@Nullable Output<String> cidrNotation) {
            $.cidrNotation = cidrNotation;
            return this;
        }

        /**
         * @param cidrNotation Address and mask in CIDR notation, e.g. &#34;147.229.15.30/31&#34;
         * 
         * @return builder
         * 
         */
        public Builder cidrNotation(String cidrNotation) {
            return cidrNotation(Output.of(cidrNotation));
        }

        /**
         * @param description Arbitrary description
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Arbitrary description
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param facility Facility where to allocate the public IP address block, makes sense only for type==public_ipv4, must be empty for type==global_ipv4, conflicts with `metro`
         * 
         * @return builder
         * 
         */
        public Builder facility(@Nullable Output<Either<String,Facility>> facility) {
            $.facility = facility;
            return this;
        }

        /**
         * @param facility Facility where to allocate the public IP address block, makes sense only for type==public_ipv4, must be empty for type==global_ipv4, conflicts with `metro`
         * 
         * @return builder
         * 
         */
        public Builder facility(Either<String,Facility> facility) {
            return facility(Output.of(facility));
        }

        /**
         * @param facility Facility where to allocate the public IP address block, makes sense only for type==public_ipv4, must be empty for type==global_ipv4, conflicts with `metro`
         * 
         * @return builder
         * 
         */
        public Builder facility(String facility) {
            return facility(Either.ofLeft(facility));
        }

        /**
         * @param facility Facility where to allocate the public IP address block, makes sense only for type==public_ipv4, must be empty for type==global_ipv4, conflicts with `metro`
         * 
         * @return builder
         * 
         */
        public Builder facility(Facility facility) {
            return facility(Either.ofRight(facility));
        }

        public Builder gateway(@Nullable Output<String> gateway) {
            $.gateway = gateway;
            return this;
        }

        public Builder gateway(String gateway) {
            return gateway(Output.of(gateway));
        }

        /**
         * @param global boolean flag whether addresses from a block are global (i.e. can be assigned in any facility)
         * 
         * @return builder
         * 
         */
        public Builder global(@Nullable Output<Boolean> global) {
            $.global = global;
            return this;
        }

        /**
         * @param global boolean flag whether addresses from a block are global (i.e. can be assigned in any facility)
         * 
         * @return builder
         * 
         */
        public Builder global(Boolean global) {
            return global(Output.of(global));
        }

        public Builder manageable(@Nullable Output<Boolean> manageable) {
            $.manageable = manageable;
            return this;
        }

        public Builder manageable(Boolean manageable) {
            return manageable(Output.of(manageable));
        }

        public Builder management(@Nullable Output<Boolean> management) {
            $.management = management;
            return this;
        }

        public Builder management(Boolean management) {
            return management(Output.of(management));
        }

        /**
         * @param metro Metro where to allocate the public IP address block, makes sense only for type==public_ipv4, must be empty for type==global_ipv4, conflicts with `facility`
         * 
         * @return builder
         * 
         */
        public Builder metro(@Nullable Output<String> metro) {
            $.metro = metro;
            return this;
        }

        /**
         * @param metro Metro where to allocate the public IP address block, makes sense only for type==public_ipv4, must be empty for type==global_ipv4, conflicts with `facility`
         * 
         * @return builder
         * 
         */
        public Builder metro(String metro) {
            return metro(Output.of(metro));
        }

        /**
         * @param netmask Mask in decimal notation, e.g. &#34;255.255.255.0&#34;
         * 
         * @return builder
         * 
         */
        public Builder netmask(@Nullable Output<String> netmask) {
            $.netmask = netmask;
            return this;
        }

        /**
         * @param netmask Mask in decimal notation, e.g. &#34;255.255.255.0&#34;
         * 
         * @return builder
         * 
         */
        public Builder netmask(String netmask) {
            return netmask(Output.of(netmask));
        }

        /**
         * @param network Network IP address portion of the block specification
         * 
         * @return builder
         * 
         */
        public Builder network(@Nullable Output<String> network) {
            $.network = network;
            return this;
        }

        /**
         * @param network Network IP address portion of the block specification
         * 
         * @return builder
         * 
         */
        public Builder network(String network) {
            return network(Output.of(network));
        }

        /**
         * @param projectId The metal project ID where to allocate the address block
         * 
         * @return builder
         * 
         */
        public Builder projectId(@Nullable Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId The metal project ID where to allocate the address block
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param public_ boolean flag whether addresses from a block are public
         * 
         * @return builder
         * 
         */
        public Builder public_(@Nullable Output<Boolean> public_) {
            $.public_ = public_;
            return this;
        }

        /**
         * @param public_ boolean flag whether addresses from a block are public
         * 
         * @return builder
         * 
         */
        public Builder public_(Boolean public_) {
            return public_(Output.of(public_));
        }

        /**
         * @param quantity The number of allocated /32 addresses, a power of 2
         * 
         * @return builder
         * 
         */
        public Builder quantity(@Nullable Output<Integer> quantity) {
            $.quantity = quantity;
            return this;
        }

        /**
         * @param quantity The number of allocated /32 addresses, a power of 2
         * 
         * @return builder
         * 
         */
        public Builder quantity(Integer quantity) {
            return quantity(Output.of(quantity));
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
         * @param type Either &#34;global_ipv4&#34; or &#34;public_ipv4&#34;, defaults to &#34;public_ipv4&#34; for backward compatibility
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable Output<Either<String,IpBlockType>> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type Either &#34;global_ipv4&#34; or &#34;public_ipv4&#34;, defaults to &#34;public_ipv4&#34; for backward compatibility
         * 
         * @return builder
         * 
         */
        public Builder type(Either<String,IpBlockType> type) {
            return type(Output.of(type));
        }

        /**
         * @param type Either &#34;global_ipv4&#34; or &#34;public_ipv4&#34;, defaults to &#34;public_ipv4&#34; for backward compatibility
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Either.ofLeft(type));
        }

        /**
         * @param type Either &#34;global_ipv4&#34; or &#34;public_ipv4&#34;, defaults to &#34;public_ipv4&#34; for backward compatibility
         * 
         * @return builder
         * 
         */
        public Builder type(IpBlockType type) {
            return type(Either.ofRight(type));
        }

        public ReservedIpBlockState build() {
            return $;
        }
    }

}
