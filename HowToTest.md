#How to Test

1. Put 20GB into S3 and Manta
2. Compute: minimal, medium, and maximum worker boundaries. Each with Go, Node and Unix.
3. Collect results

##Regarding 1

How is it possible to upload 20GB (40GB total) to S3 and Manta?

Possible best option: Spin up 40 Machines, and let each Machine write 0.5GB to S3 and Manta using their command line tools:

```
//e.g.

sdc-createmachine -n getting-started -e 3390ca7c-f2e7-11e1-8818-c36e0b12e58b
ssh-add
ssh -A admin@10.88.88.50 '/scripts/loggeneration.sh'

sleep 2h

sdc-stopmachine 3390ca7c-f2e7-11e1-8818-c36e0b12e58b
dc-deletemachine 3390ca7c-f2e7-11e1-8818-c36e0b12e58b

```
