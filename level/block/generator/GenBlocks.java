package pers.tnze.gomc.gen;

import com.google.common.collect.ImmutableMap;
import net.minecraft.SharedConstants;
import net.minecraft.core.registries.BuiltInRegistries;
import net.minecraft.nbt.*;
import net.minecraft.resources.ResourceLocation;
import net.minecraft.server.Bootstrap;
import net.minecraft.world.level.block.Block;
import net.minecraft.world.level.block.Blocks;
import net.minecraft.world.level.block.entity.BlockEntityType;
import net.minecraft.world.level.block.state.BlockState;
import net.minecraft.world.level.block.state.properties.EnumProperty;
import net.minecraft.world.level.block.state.properties.Property;

import java.io.DataOutput;
import java.io.DataOutputStream;
import java.io.FileOutputStream;
import java.util.Map;
import java.util.Objects;
import java.util.zip.GZIPOutputStream;

public class GenBlocks {

    public static void main(String[] args) throws Exception {
        System.out.println("program start!");
        SharedConstants.tryDetectVersion();
        Bootstrap.bootStrap();
        Blocks.rebuildCache();

        try (FileOutputStream f = new FileOutputStream("blocks.nbt")) {
            try (GZIPOutputStream g = new GZIPOutputStream(f)) {
                DataOutput writer = new DataOutputStream(g);
                NbtIo.writeUnnamedTag(getBlocksWithMeta(), writer);
            }
        }
        try (FileOutputStream f = new FileOutputStream("block_states.nbt")) {
            try (GZIPOutputStream g = new GZIPOutputStream(f)) {
                DataOutput writer = new DataOutputStream(g);
                NbtIo.writeUnnamedTag(getBlockStates(), writer);
            }
        }
        try (FileOutputStream f = new FileOutputStream("block_entities.nbt")) {
            try (GZIPOutputStream g = new GZIPOutputStream(f)) {
                DataOutput writer = new DataOutputStream(g);
                NbtIo.writeUnnamedTag(genBlockEntities(), writer);
            }
        }
    }

    private static ListTag getBlocksWithMeta() throws Exception {
        ListTag list = new ListTag();
        for (Block block : BuiltInRegistries.BLOCK) {
            BlockState state = block.defaultBlockState();
            CompoundTag b = new CompoundTag();
            b.putString("Name", BuiltInRegistries.BLOCK.getKey(block).toString());
            ImmutableMap<Property<?>, Comparable<?>> values = state.getValues();
            if (!values.isEmpty()) {
                CompoundTag meta = new CompoundTag();
                for (Map.Entry<Property<?>, Comparable<?>> entry : values.entrySet()) {
                    Property<?> key = entry.getKey();
                    Comparable<?> value = entry.getValue();
                    String name = key.getName();
                    String typeName;
                    if (key instanceof EnumProperty<?>) {
                        if (value.getClass().getName().contains("net.minecraft.core.Direction$Axis")) {
                            typeName = "Axis";
                        } else {
                            typeName = value.getClass().getSimpleName();
                        }
                        if (typeName.isBlank()) {
                            throw new Exception("Type is blank: " + value.getClass().getName());
                        }
                    } else {
                        typeName = key.getClass().getSimpleName();
                    }
                    meta.putString(name, typeName);
                }
                b.put("Meta", meta);
            }
            list.add(b);
        }
        return list;
    }

    private static ListTag getBlockStates() {
        ListTag list = new ListTag();
        for (BlockState blockState : Block.BLOCK_STATE_REGISTRY) {
            list.add(NbtUtils.writeBlockState(blockState));
        }
        return list;
    }

    private static ListTag genBlockEntities() {
        ListTag list = new ListTag();
        for (BlockEntityType<?> blockEntity : BuiltInRegistries.BLOCK_ENTITY_TYPE) {
            ResourceLocation value = BuiltInRegistries.BLOCK_ENTITY_TYPE.getKey(blockEntity);
            ListTag validBlocksList = new ListTag();
            for (Block validBlock : blockEntity.validBlocks){
                validBlocksList.add(StringTag.valueOf(BuiltInRegistries.BLOCK.getKey(validBlock).toString()));
            }
            CompoundTag be = new CompoundTag();
            be.putString("Name", Objects.requireNonNull(value).toString());
            be.put("ValidBlocks", validBlocksList);

            list.add(be);
        }
        return list;
    }
}
