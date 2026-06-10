import Fastify from 'fastify';
import cors from '@fastify/cors';

const app = Fastify({ logger: true });

await app.register(cors);

app.get('/health', async () => ({
  status: 'healthy',
  service: 'messaging',
}));

app.get('/messages', async () => ({ messages: [] }));

app.post('/messages', async (request, reply) => {
  reply.code(201);
  return { id: 'new-message', status: 'sent' };
});

app.get('/messages/:id', async () => ({ id: 'message-id', content: 'hello' }));

app.post('/messages/:id/read', async () => ({ status: 'read' }));

const start = async () => {
  try {
    await app.listen({ port: 8088, host: '0.0.0.0' });
    app.log.info('Messaging service listening on :8088');
  } catch (err) {
    app.log.error(err);
    process.exit(1);
  }
};

start();
