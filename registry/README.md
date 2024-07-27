# Registry System

instance {
    ResourceLocation {
        namespace: string
        path: string
    }

    ResourceKey<T> {
        registryName: ResourceLocation
        location: ResourceLocation
    }

    registry<T> {
        getId(T): int
        byId(int): T

        getKey(T): ResourceLocation
        getResourceKey(T): ResourceKey<T>
        get(ResourceKey<T>): T
        get(ResourceLocation): T

        getTags(TagKey<T>): (TagKey<T>, *T[])[]
        getTagNames(): TagKey<T>[]
        resetTags()
        bindTags((TagKey<T>, *T[])[])
    }[]

    TagKey<T> {
        *Registry<T>
        ResourceLocation
    }
}
