package pers.tnze.gomc.gen;

import it.unimi.dsi.fastutil.ints.Int2ObjectMap;
import net.minecraft.network.ConnectionProtocol;
import net.minecraft.network.protocol.Packet;
import net.minecraft.network.protocol.PacketFlow;

import java.io.FileWriter;
import java.io.IOException;
import java.io.Writer;

public class Main {

    public static void main(String[] args) throws Exception {
        try (FileWriter w = new FileWriter("packet_names.txt")) {
            handlePackets(w, ConnectionProtocol.LOGIN.getPacketsByIds(PacketFlow.CLIENTBOUND));
            System.out.println();
            handlePackets(w, ConnectionProtocol.LOGIN.getPacketsByIds(PacketFlow.SERVERBOUND));
            System.out.println();
            System.out.println();
            handlePackets(w, ConnectionProtocol.STATUS.getPacketsByIds(PacketFlow.CLIENTBOUND));
            System.out.println();
            handlePackets(w, ConnectionProtocol.STATUS.getPacketsByIds(PacketFlow.SERVERBOUND));
            System.out.println();
            System.out.println();
            handlePackets(w, ConnectionProtocol.PLAY.getPacketsByIds(PacketFlow.CLIENTBOUND));
            System.out.println();
            handlePackets(w, ConnectionProtocol.PLAY.getPacketsByIds(PacketFlow.SERVERBOUND));
        }
    }

    private static void handlePackets(Writer w, Int2ObjectMap<Class<? extends Packet<?>>> packets) throws IOException {
        for (int i = 0; i < packets.size(); i++) {
            Class<? extends Packet<?>> c = packets.get(i);
            String className = c.getSimpleName();
            if (className.endsWith("Packet"))
                className = className.substring(0, className.length() - "Packet".length());
            else {
                String superClassName = c.getSuperclass().getSimpleName();
                if (superClassName.endsWith("Packet"))
                    className = superClassName.substring(0, superClassName.length() - "Packet".length()) + className;
            }
            System.out.println(className);
            w.write(className + "\n");
        }
        w.write('\n');
    }
}