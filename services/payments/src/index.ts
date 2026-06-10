import Fastify from 'fastify';
import cors from '@fastify/cors';

const app = Fastify({ logger: true });

await app.register(cors, { origin: true });

app.get('/healthz', async () => ({ status: 'ok' }));
app.get('/readyz', async () => ({ status: 'ready' }));

app.get('/api/v1/payments', async () => {
  return { payments: [], total: 0 };
});

app.post('/api/v1/payments', async (request, reply) => {
  const body = request.body as any;
  return reply.status(201).send({
    payment: {
      id: crypto.randomUUID(),
      ...body,
      status: 'pending',
      created_at: new Date().toISOString(),
    },
  });
});

app.get('/api/v1/payments/:id', async (request, reply) => {
  const { id } = request.params as any;
  return {
    payment: {
      id,
      status: 'completed',
      amount: 100.00,
      currency: 'USD',
    },
  };
});

app.post('/api/v1/payments/:id/process', async (request, reply) => {
  const { id } = request.params as any;
  return {
    payment: {
      id,
      status: 'completed',
      provider_ref: `PAY-${id.substring(0, 8)}`,
    },
  };
});

app.post('/api/v1/payments/:id/refund', async (request, reply) => {
  const { id } = request.params as any;
  return { message: 'refund processed', payment_id: id };
});

const start = async () => {
  const port = parseInt(process.env.PORT || '8087');
  await app.listen({ port, host: '0.0.0.0' });
  console.log(`Payments service starting on port ${port}`);
};

start();
