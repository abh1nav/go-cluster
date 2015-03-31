#Go-Cluster-Test

Basic clustering in go.

## How To

1. Clone this repo.  
  
2. `make build`  
  
3. Start the seed node  

```
./bin/go-cluster-0.1.0.bin --node="hodor"
```

It'll look like:

![Seed Node](https://raw.github.com/abh1nav/go-cluster-test/master/img/seed_node.png)

4. Start another node and give it the seed's location  
  
```
./bin/go-cluster-0.1.0.bin --bind="127.0.0.1:5001" \
	--node="john-snow" --seed="127.0.0.1:6000" \
	--advertise-addr="127.0.0.1" --advertise-port=6001
```

It'll find the seed and output:

![Cluster Member](https://raw.github.com/abh1nav/go-cluster-test/master/img/cluster_member.png)

The seed will then show:

![Seed Reaction](https://raw.github.com/abh1nav/go-cluster-test/master/img/seed_reaction.png)