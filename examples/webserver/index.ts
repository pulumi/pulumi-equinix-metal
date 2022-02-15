import * as pulumi from "@pulumi/pulumi";
import * as metal from "@pulumi/equinix-metal";

const project = metal.getProject({ name: "ci-project"}).then( x=> x.id);

const vm = new metal.Device("vm", {
    metro: "da",
    billingCycle: metal.BillingCycle.Hourly,
    hostname: "demoinstance",
    operatingSystem: metal.OperatingSystem.Ubuntu2004,
    plan: metal.Plan.C3SmallX86,
    projectId: project,
});

export const ip = vm.accessPublicIpv4;
