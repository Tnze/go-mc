// This file is used to generate blocks.nbt and block_states.nbt
// The necessary environment can be generated from https://github.com/Hexeption/MCP-Reborn
package mcp.client

import net.minecraft.SharedConstants
import net.minecraft.core.Direction
import net.minecraft.core.Direction.Axis
import net.minecraft.core.Registry
import net.minecraft.nbt.CompoundTag
import net.minecraft.nbt.ListTag
import net.minecraft.nbt.NbtIo
import net.minecraft.nbt.NbtUtils
import net.minecraft.server.Bootstrap
import net.minecraft.world.level.block.Block
import net.minecraft.world.level.block.Blocks
import net.minecraft.world.level.block.grower.SpruceTreeGrower
import net.minecraft.world.level.block.state.BlockBehaviour.Properties
import net.minecraft.world.level.block.state.properties.AttachFace
import net.minecraft.world.level.block.state.properties.BambooLeaves
import net.minecraft.world.level.block.state.properties.BlockStateProperties
import net.minecraft.world.level.block.state.properties.BooleanProperty
import net.minecraft.world.level.block.state.properties.EnumProperty
import net.minecraft.world.level.block.state.properties.IntegerProperty
import net.minecraft.world.level.block.state.properties.NoteBlockInstrument
import net.minecraft.world.level.block.state.properties.Property
import net.minecraft.world.level.block.state.properties.RailShape
import net.minecraft.world.level.material.Fluid
import net.minecraft.world.level.material.Fluids
import java.io.DataOutputStream
import java.io.FileOutputStream
import java.io.ObjectOutputStream
import java.lang.reflect.Modifier
import java.util.zip.GZIPOutputStream
import kotlin.reflect.KClass

object Start {
    @JvmStatic
    fun main(args: Array<String>) {
        println("program start!")
        SharedConstants.tryDetectVersion()
        Bootstrap.bootStrap()
        Blocks.rebuildCache()

        FileOutputStream("blocks.nbt").use { stream ->
            GZIPOutputStream(stream).use {
                NbtIo.writeUnnamedTag(blocksWithMeta(), DataOutputStream(it))
            }
        }

        FileOutputStream("block_states.nbt").use { stream ->
            GZIPOutputStream(stream).use {
                NbtIo.writeUnnamedTag(blockStates(), DataOutputStream(it))
            }
        }

        FileOutputStream("fluid_states.nbt").use { stream ->
            GZIPOutputStream(stream).use {
                NbtIo.writeUnnamedTag(fluidStates(), DataOutputStream(it))
            }
        }
    }

    private fun blocksWithMeta(): ListTag {
        return ListTag().apply {
            Registry.BLOCK.forEach { block ->
                val states = CompoundTag().apply {
                    putString("Name", Registry.BLOCK.getKey(block).toString())
                    val meta = CompoundTag().apply {
                        block.defaultBlockState().values.forEach { (key, value) ->
                            val name: String = key.name
                            val typeName = if (key is EnumProperty<*> && value.javaClass.name.contains("net.minecraft.core.Direction\$Axis")) {
                                "Axis"
                            } else {
                                value.javaClass.simpleName
                            }
                            if (typeName.isBlank()) {
                                throw java.lang.Exception("Type is blank: " + value.javaClass.name)
                            }
                            putString(name, typeName)
                        }
                    }

                    val properties = CompoundTag().apply {
                        putBoolean("HasCollision", block.hasCollision)
                        putFloat("ExplosionResistance", block.explosionResistance)
                        putFloat("DestroyTime", block.destroyTime)
                        putBoolean("RequiresCorrectToolForDrop", block.requiresCorrectToolForDrops)
                        putFloat("Friction", block.friction)
                        putFloat("SpeedFactor", block.speedFactor)
                        putFloat("JumpFactor", block.jumpFactor)
                        putBoolean("CanOcclude", block.canOcclude)
                        putBoolean("IsAir", block.isAir)
                        putBoolean("DynamicShape", block.dynamicShape)
                    }

                    val defaultValues = CompoundTag().apply {
                        block.defaultBlockState().values.forEach { (key, value) ->
                            when {
                                key is BooleanProperty && value is Boolean -> putBoolean(key.name.capitalize(), value)
                                key is IntegerProperty && value is Int -> putInt(key.name.capitalize(), (key.min shl 16) or (key.max and 0xFFFF))
                                key is EnumProperty && value is Enum -> putString(key.name.capitalize(), toGoTypeName(value.javaClass, value.name))
                                else -> println("Unknown type: " + value.javaClass.name)
                            }
                        }
                    }

                    put("Properties", properties)
                    put("Meta", meta)
                    put("Default", defaultValues)
                }

                // Put the data into the nbt
                add(states)
            }
        }
    }

    private fun blockStates(): ListTag = ListTag().apply { addAll(Block.BLOCK_STATE_REGISTRY.map { NbtUtils.writeBlockState(it) }) }

    private fun fluidStates(): ListTag = ListTag().apply { addAll(Fluid.FLUID_STATE_REGISTRY.map { NbtUtils.writeFluidState(it) }) }


    private fun toGoTypeName(clazz: Class<*>, str: String): String {
        // Example: class: WallSide, str: NONE -> WallSideNone
        return clazz.simpleName + str.lowercase().capitalize()
    }
}
