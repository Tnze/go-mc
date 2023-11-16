## Updating `data`

1. Go to [https://github.com/PrismarineJS/minecraft-data/tree/master/data/pc/{version}](https://github.com/PrismarineJS/minecraft-data/tree/master/data/pc)
2. Update `version` in the following files if there is a new corresponding JSON file available:
   - [gen_block.go](block/gen_block.go) - `blocks.json`
   - [gen_entity.go](entity/gen_entity.go) - `entities.json`
   - [gen_item.go](item/gen_item.go) - `items.json`
3. Update the `URL` in [gen_soundid.go](soundid/gen_soundid.go) (verify the URL returns a response first)
4. Run `go generate ./...`