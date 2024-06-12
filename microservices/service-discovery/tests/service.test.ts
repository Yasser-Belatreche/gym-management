import assert from 'node:assert';
import { after, before, beforeEach, describe, it } from 'node:test';

import { Redis } from '../src/redis.js';
import { ServiceDiscovery } from '../src/service.js';

await describe('Service Discovery', async () => {
    let redis: Redis;
    let service: ServiceDiscovery;

    before(async () => {
        redis = Redis.Instance({ url: process.env.REDIS_URL! });
        await redis.connect();

        service = new ServiceDiscovery(redis.Client());
    });

    after(async () => {
        await redis.Client().disconnect();
    });

    beforeEach(async () => {
        await redis.Client().flushAll();
    });

    await it('should register a service', async () => {
        await service.register('test', 'http://localhost:3000');

        const s = await service.getService('test');

        assert.strictEqual(s?.name, 'test');
        assert.strictEqual(s?.url, 'http://localhost:3000');
    });

    await it('should be able to update a service', async () => {
        await service.register('test', 'http://localhost:3000');
        await service.register('test', 'http://localhost:3001');

        const s = await service.getService('test');

        assert.strictEqual(s?.name, 'test');
        assert.strictEqual(s?.url, 'http://localhost:3001');
    });

    await it('should get all services', async () => {
        await service.register('test', 'http://localhost:3000');
        await service.register('test2', 'http://localhost:3001');

        const services = await service.getServices();

        assert.strictEqual(services.length, 2);
    });

    await it('should delete a service', async () => {
        await service.register('test', 'http://localhost:3000');
        await service.deleteService('test');

        const s = await service.getService('test');

        assert.strictEqual(s, undefined);
    });
});
