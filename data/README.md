## Updating `data`

1. Go to [https://github.com/PrismarineJS/minecraft-data/tree/master/data/pc/{version}](https://github.com/PrismarineJS/minecraft-data/tree/master/data/pc)
2. Update the URL (if appropriate) in [gen_block.go](block/gen_block.go), [gen_entity.go](entity/gen_entity.go),
[gen_item.go](item/gen_item.go), and [gen_packetid.go](packetid/gen_packetid.go)
3. Update the `URL` in [gen_soundid.go](soundid/gen_soundid.go)
4. Run `go generate ./...`