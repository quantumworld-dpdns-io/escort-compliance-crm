import Fastify from 'fastify';
import cors from '@fastify/cors';

const app = Fastify({ logger: true });

await app.register(cors);

app.get('/health', async () => ({
  status: 'healthy',
  service: 'payments',
}));

app.get('/payments', async () => ({ payments: [] }));

app.post('/payments', async (request, reply) => {
  reply.code(201);
  return { id: 'new-payment', status: 'created' };
});

app.get('/payments/:id', async () => ({ id: 'payment-id', status: 'pending' }));

app.post('/payments/:id/process', async () => ({ status: 'processed' }));

app.post('/payments/:id/refund', async () => ({ status: 'refunded' }));

const start = async () => {
  try {
    await app.listen({ port: 8087, host: '0.0.0.0' });
    app.log.info('Payments service listening on :8087');
  } catch (err) {
    app.log.error(err);
    process.exit(1);
  }
};

start();
