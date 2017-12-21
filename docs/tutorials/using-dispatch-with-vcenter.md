---
layout: page
title:  "Tutorial: Using Dispatch with vCenter"
---
[Dispatch Home]({{site.baseurl}}/) &raquo; Tutorial: Using Dispatch with vCenter
# Tutorial: Using Dispatch with vCenter

## Description
This tutorial will guide you to write a serverless function that can be executed in response to events from your vCenter server. 

## Overview
Serverless functions in Dispatch can be automatically executed based on events from a specific source or manually executed by the user. For Dispatch to monitor events from a source, an Event Driver must exist for that source. By default, Dispatch provides certain Event Drivers e.g for a VMware vCenter Server. In this tutorial, you will learn how to create an Event Driver instance for your vCenter Server and write a serverless function that will post messages to a Slack channel whenever a VM is created in vCenter.

Following is a summary of the serverless workflow :-

1. User creates a VM in vCenter.

1. Dispatch's vCenter Event Driver detects the VmBeingCreatedEvent event.

1. Dispatch's Event Manager triggers the subscribed serverless function.

1. The function posts a slack message on a specific channel with the details of the created VM.

## Prerequisites
* Dispatch framework must be installed

* VMware vCenter Server

* Slack Account with privileges to create an incoming webhook

## Steps

### Step 1: Create an instance of the vCenter Event Driver in Dispatch
```
dispatch create event-driver <name> vcenter --set vcenterurl='<user>:<password>@<vcenter_host>:443'
```
where
<dl>
<dt>name</dt>
<dd>The name of the event driver instance e.g my-vcenter</dd>
<dt>user</dt>
<dd>The vCenter user e.g administrator@vsphere.local</dd>
<dt>password</dt>
<dd>The vCenter password</dd>
<dt>vcenter_host</dt>
<dd>The vCenter Server Host IP address or hostname</dd>
</dl>

### Step 2: Configure an incoming webhook in your Slack Account
Set up an incoming webhook integration in your slack workspace by following the instructions in this <a href="https://api.slack.com/incoming-webhooks" target="_blank">page</a>. Once you have setup the integration, make a note of the webhook URL.

Create a json file `secret.json` with the incoming webhook URL e.g.
```
{
    "slack_url": "https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX"
}
```

Create a dispatch secret to store your Slack webhook URL.
```
dispatch create secret slack secret.json
```

### Step 3: Create a sample serverless function to post messages to Slack

Create the base image required to execute the sample serverless function that will post messages to Slack.
```
dispatch create base-image node-request vmware/openfaas-nodejs-with-request:0.0.1-dev1 --language nodejs6 --public
```
Download and create the sample serverless function.
```
curl -LO https://raw.githubusercontent.com/vmware/dispatch/master/examples/nodejs6/slack_vm_being_deployed.js

dispatch create function node-request slack-post-message slack_vm_being_deployed.js
```

### Step 4: Subscribe to the vCenter Event
Subscribe to the event **vm.being.deployed** that is published by the vCenter Event Driver and specify the name of the function **slack-post-message** that must be executed when the event occurs.
```
dispatch create subscription vm.being.deployed slack-post-message --secret slack
```

### Step 5: Create a VM in your vCenter Server
Create a VM in your vCenter Server to watch the serverless function get executed and post a message to your Slack channel.
