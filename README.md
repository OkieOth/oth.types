some example to use yacg to generate golang types

# Usage

```bash
# install dependency to the project
go get github.com/google/uuid

# generate random data for the models
docker run -u $(id -u ${USER}):$(id -g ${USER}) -v `pwd`:/resources --rm -t --entrypoint "python3" ghcr.io/okieoth/yacg:6.1.0 createRandomData.py --model /resources/configs/models/model_0101.json --defaultElemCount 10  --outputDir /resources/configs/models/example_data/0101 --defaultTypeDepth 10 --allTypes

docker run -u $(id -u ${USER}):$(id -g ${USER}) -v `pwd`:/resources --rm -t --entrypoint "python3" ghcr.io/okieoth/yacg:6.1.0 createRandomData.py --model /resources/configs/models/model_0102.json --defaultElemCount 10  --outputDir /resources/configs/models/example_data/0102 --defaultTypeDepth 10 --allTypes

docker run -u $(id -u ${USER}):$(id -g ${USER}) -v `pwd`:/resources --rm -t --entrypoint "python3" ghcr.io/okieoth/yacg:6.1.0 createRandomData.py --model /resources/configs/models/model_0201.json --defaultElemCount 10  --outputDir /resources/configs/models/example_data/0201 --defaultTypeDepth 10 --allTypes

docker run -u $(id -u ${USER}):$(id -g ${USER}) -v `pwd`:/resources --rm -t --entrypoint "python3" ghcr.io/okieoth/yacg:6.1.0 createRandomData.py --model /resources/configs/models/model_0202.json --defaultElemCount 10  --outputDir /resources/configs/models/example_data/0202 --defaultTypeDepth 10 --allTypes
```