import Fastify from 'fastify';
import cors from '@fastify/cors';
import { v4 as uuid } from 'uuid';

const app = Fastify({ logger: true });

await app.register(cors, { origin: true });

interface Message {
  id: string;
  conversation_id: string;
  sender_id: string;
  recipient_id: string;
  content: string;
  content_type: string;
  is_read: boolean;
  created_at: string;
}

const messages = new Map<string, Message>();
const conversations = new Map<string, string[]>();

app.get('/healthz', async () => ({ status: 'ok' }));
app.get('/readyz', async () => ({ status: 'ready' }));

app.post('/api/v1/messages', async (request, reply) => {
  const { conversation_id, sender_id, recipient_id, content, content_type } = request.body as any;
  
  const msg: Message = {
    id: uuid(),
    conversation_id: conversation_id || uuid(),
    sender_id,
    recipient_id,
    content,
    content_type: content_type || 'text',
    is_read: false,
    created_at: new Date().toISOString(),
  };
  
  messages.set(msg.id, msg);
  
  if (!conversations.has(msg.conversation_id)) {
    conversations.set(msg.conversation_id, []);
  }
  conversations.get(msg.conversation_id)!.push(msg.id);
  
  return reply.status(201).send({ message: msg });
});

app.get('/api/v1/messages/:conversation_id', async (request, reply) => {
  const { conversation_id } = request.params as any;
  const msgIds = conversations.get(conversation_id) || [];
  const msgs = msgIds.map(id => messages.get(id)).filter(Boolean);
  return { messages: msgs };
});

app.post('/api/v1/messages/:id/read', async (request, reply) => {
  const { id } = request.params as any;
  const msg = messages.get(id);
  if (!msg) return reply.status(404).send({ error: 'message not found' });
  msg.is_read = true;
  return { message: msg };
});

app.get('/api/v1/conversations/:user_id', async (request, reply) => {
  const { user_id } = request.params as any;
  const userConversations = [];
  for (const [convId, msgIds] of conversations.entries()) {
    const lastMsgId = msgIds[msgIds.length - 1];
    const lastMsg = messages.get(lastMsgId);
    if (lastMsg && (lastMsg.sender_id === user_id || lastMsg.recipient_id === user_id)) {
      userConversations.push({ id: convId, last_message: lastMsg });
    }
  }
  return { conversations: userConversations };
});

const start = async () => {
  const port = parseInt(process.env.PORT || '8088');
  await app.listen({ port, host: '0.0.0.0' });
  console.log(`Messaging service starting on port ${port}`);
};

start();
