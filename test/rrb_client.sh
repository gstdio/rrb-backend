curl http://localhost:10001/rrbend/class/getall
curl http://localhost:10001/rrbend/subclass/getall
curl http://localhost:10001/rrbend/subclass?class_id=1
curl http://localhost:10001/rrbend/product/subclass?id=1
curl http://localhost:10001/rrbend/product/shop?id=1

curl http://localhost:10001/rrbend/class/insert -d '{"name": "class2", "desc": "a class2"}'
curl http://localhost:10001/rrbend/class/update -d '{"id": 3, "name": "class2", "desc": "a class22"}'
curl http://localhost:10001/rrbend/subclass/insert -d '{"class_id: "3", "name": "subclass2", "desc": "a subclass2"}'
curl http://localhost:10001/rrbend/subclass/update -d '{"id": 3, "class_id": 3, "name": "subclass2", "desc": "a subclass22"}'
curl http://localhost:10001/rrbend/product/insert -d '{"subclass_id": 1, "name": "orange", "desc": "an orange"}'
curl http://localhost:10001/rrbend/product/update -d '{"id": 3, "subclass_id": 1, "name": "orange", "desc": "an orange2"}'
