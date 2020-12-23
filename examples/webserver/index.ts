import * as pulumi from "@pulumi/pulumi";
import * as metal from "@pulumi/equinix-metal";

const project = metal.getProject({ name: "ci-project"}).then( x=> x.id);

const vm = new metal.Device("vm", {
    facilities: [metal.Facility.EWR1],
    billingCycle: metal.BillingCycle.Hourly,
    hostname: "demoinstance",
    operatingSystem: metal.OperatingSystem.Ubuntu1804,
    plan: metal.Plan.T1SmallX86,
    projectId: project,
});

export const ip = vm.accessPublicIpv4;
