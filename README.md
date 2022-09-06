# Posture Tracker

A simple posture tracking web service.

## Requirements
- Golang
- Ansible
- Raspberry Pi (Should work with another Linux like systems)

If you are not deploying to a Raspberry Pi change the build parameters in the `Makefile` to suite your needs.

## Deployment

Create a file called `hosts` like this.

```
host1
```

This is an [Ansible inventory file](https://docs.ansible.com/ansible/latest/user_guide/intro_inventory.html).

To deploy run `make deploy clean`. This will copy the files needed to run the service and start it up.
