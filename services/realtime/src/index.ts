import Fastify from 'fastify';
import cors from '@fastify/cors';
import websocket from '@fastify/websocket';

const app = Fastify({ logger: true });

await app.register(cors);
await app.register(websocket);

app.get('/health', async () => ({
  status: 'healthy',
  service: 'realtime',
}));

app.get('/ws', { websocket: true }, (socket, req) => {
  app.log.info('WebSocket client connected');

  socket.on('message', (message) => {
    try {
      const data = JSON.parse(message.toString());
      socket.send(JSON.stringify({ type: 'ack', data }));
    } catch (e) {
      socket.send(JSON.stringify({ type: 'error', message: 'Invalid JSON' }));
    }
  });

  socket.on('close', () => {
    app.log.info('WebSocket client disconnected');
  });
});

const start = async () => {
  try {
    await app.listen({ port: 8089, host: '0.0.0.0' });
    app.log.info('Realtime service listening on :8089');
  } catch (err) {
    app.log.error(err);
    process.exit(1);
  }
};

start();
