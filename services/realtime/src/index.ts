import Fastify from 'fastify';
import cors from '@fastify/cors';
import websocket from '@fastify/websocket';

const app = Fastify({ logger: true });

await app.register(cors, { origin: true });
await app.register(websocket);

const clients = new Map<string, any>();
const channels = new Map<string, Set<string>>();

app.get('/healthz', async () => ({ status: 'ok' }));
app.get('/readyz', async () => ({ status: 'ready' }));

app.get('/ws', { websocket: true }, (socket, request) => {
  const clientId = Math.random().toString(36).substring(7);
  clients.set(clientId, socket);
  
  console.log(`Client connected: ${clientId}`);
  
  socket.on('message', (data) => {
    try {
      const msg = JSON.parse(data.toString());
      
      switch (msg.type) {
        case 'subscribe':
          if (!channels.has(msg.channel)) {
            channels.set(msg.channel, new Set());
          }
          channels.get(msg.channel)!.add(clientId);
          socket.send(JSON.stringify({ type: 'subscribed', channel: msg.channel }));
          break;
          
        case 'unsubscribe':
          channels.get(msg.channel)?.delete(clientId);
          socket.send(JSON.stringify({ type: 'unsubscribed', channel: msg.channel }));
          break;
          
        case 'publish':
          const channel = channels.get(msg.channel);
          if (channel) {
            for (const id of channel) {
              const client = clients.get(id);
              if (client && id !== clientId) {
                client.send(JSON.stringify({
                  type: 'message',
                  channel: msg.channel,
                  data: msg.data,
                  timestamp: new Date().toISOString(),
                }));
              }
            }
          }
          break;
          
        case 'ping':
          socket.send(JSON.stringify({ type: 'pong', timestamp: new Date().toISOString() }));
          break;
      }
    } catch (err) {
      console.error('Invalid message:', err);
    }
  });
  
  socket.on('close', () => {
    clients.delete(clientId);
    for (const [, channel] of channels.entries()) {
      channel.delete(clientId);
    }
    console.log(`Client disconnected: ${clientId}`);
  });
});

app.get('/api/v1/realtime/stats', async () => {
  return {
    connected_clients: clients.size,
    channels: Object.fromEntries(
      Array.from(channels.entries()).map(([name, members]) => [name, members.size])
    ),
  };
});

const start = async () => {
  const port = parseInt(process.env.PORT || '8089');
  await app.listen({ port, host: '0.0.0.0' });
  console.log(`Realtime service starting on port ${port}`);
};

start();
