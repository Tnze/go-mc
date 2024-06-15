package pers.tnze.gomc.gen;

import net.minecraft.SharedConstants;
import net.minecraft.network.ProtocolInfo;
import net.minecraft.network.protocol.configuration.ConfigurationProtocols;
import net.minecraft.network.protocol.game.GameProtocols;
import net.minecraft.network.protocol.login.LoginProtocols;
import net.minecraft.network.protocol.status.StatusProtocols;
import net.minecraft.server.Bootstrap;

import java.io.FileWriter;
import java.io.IOException;
import java.io.Writer;

public class GenPacketId {

    public static void main(String[] args) throws Exception {
        SharedConstants.tryDetectVersion();
        Bootstrap.bootStrap();
        try (FileWriter w = new FileWriter("packet_names.txt")) {
            handlePackets(w, LoginProtocols.CLIENTBOUND_TEMPLATE, "ClientboundLogin");
            handlePackets(w, LoginProtocols.SERVERBOUND_TEMPLATE, "ServerboundLogin");
            handlePackets(w, StatusProtocols.CLIENTBOUND_TEMPLATE, "ClientboundStatus");
            handlePackets(w, StatusProtocols.SERVERBOUND_TEMPLATE, "ServerboundStatus");
            handlePackets(w, ConfigurationProtocols.CLIENTBOUND_TEMPLATE, "ClientboundConfig");
            handlePackets(w, ConfigurationProtocols.SERVERBOUND_TEMPLATE, "ServerboundConfig");
            handlePackets(w, GameProtocols.CLIENTBOUND_TEMPLATE, "");
            handlePackets(w, GameProtocols.SERVERBOUND_TEMPLATE, "");
        }
    }

    private static void handlePackets(Writer w, ProtocolInfo.Unbound<?, ?> packets, String prefix) throws IOException {
        packets.listPackets((packetType, i) -> {
            String packetName = packetType.id().getPath();
            String[] words = packetName.split("_");
            try {
                if (prefix != null) {
                    w.write(prefix);
                }
                for (String word : words) {
                    w.write(Character.toUpperCase(word.charAt(0)));
                    w.write(word.substring(1));
                }
                w.write("\n");
            } catch (IOException e) {
                throw new RuntimeException(e);
            }
        });
        w.write('\n');
    }
}