## Deployment

The PoC can be deployed to a single GCP compute instance with the Container Optimized OS base image. The deployed application currently serves some supporting functions for [trav.finsyn.se](http://trav.finsyn.se/systems/V64_2019-09-26_26_4.html). Beware that a deploy will result in some downtime.

```
bigt@torlap1:~/travtech/api$ ./deploy/deploy.sh 
docker-compose.yml                                                                                                                                           100%  602    35.8KB/s   00:00    
api-manager.sh                                                                                                                                               100%  375    26.1KB/s   00:00    
latest: Pulling from travtech/faas-gateway
Digest: sha256:d8923364d0713f21eb0226c512c369d22dd673714d42838d639ad51be0367db6
Status: Image is up to date for eu.gcr.io/travtech/faas-gateway:latest
eu.gcr.io/travtech/faas-gateway:latest
latest: Pulling from travtech/rhymer
Digest: sha256:b51a0bd1b853ddb61afa98fb8cdbc0cbb09ac995751fb293ec56775c62a372f7
Status: Image is up to date for eu.gcr.io/travtech/rhymer:latest
eu.gcr.io/travtech/rhymer:latest
latest: Pulling from travtech/reducer
Digest: sha256:761a31cd8eefdafddf302c34339733488fd2b912269023eecae2406b4badd993
Status: Image is up to date for eu.gcr.io/travtech/reducer:latest
eu.gcr.io/travtech/reducer:latest
Stopping bigt_rhymer_1       ... done
Stopping bigt_reducer_1      ... done
Stopping bigt_faas-gateway_1 ... done
Removing bigt_rhymer_1       ... done
Removing bigt_reducer_1      ... done
Removing bigt_faas-gateway_1 ... done
Removing network bigt_default
Creating network "bigt_default" with the default driver
Creating bigt_reducer_1      ... done
Creating bigt_faas-gateway_1 ... done
Creating bigt_rhymer_1       ... done
Attaching to bigt_rhymer_1, bigt_faas-gateway_1, bigt_reducer_1
rhymer_1        | rhymer initiating
faas-gateway_1  | gateway initiating
faas-gateway_1  | service discovery initiated
faas-gateway_1  | proxy server started
faas-gateway_1  | service discovered: /rhymer
faas-gateway_1  | service discovered: /reducer
```
