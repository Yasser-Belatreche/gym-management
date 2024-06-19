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
        const { id } = await service.register('test', 'http://localhost:3000');

        const s = await service.getService(id);

        assert.strictEqual(s?.name, 'test');
        assert.strictEqual(s?.url, 'http://localhost:3000');
    });

    await it('should give each service a unique id', async () => {
        const { id: id1 } = await service.register('test', 'http://localhost:3000');
        const { id: id2 } = await service.register('test', 'http://localhost:3001');

        assert.notStrictEqual(id1, id2);
    });

    await it('should not be able to register a service with an invalid url', async () => {
        try {
            await service.register('test', 'localhost:3000');
            assert.fail('Should have thrown an error');
        } catch (error) {
            assert.ok(error instanceof Error);
        }
    });

    await it('should be able to update a service by id', async () => {
        const { id } = await service.register('test', 'http://localhost:3000');

        await service.updateService(id, 'newName', 'http://localhost:3001');

        const s = await service.getService(id);

        assert.strictEqual(s?.name, 'newName');
        assert.strictEqual(s?.url, 'http://localhost:3001');
    });

    await it("should not be able to update a service that doesn't exist", async () => {
        try {
            await service.updateService('invalidId', 'newName', 'http://localhost:3000');
            assert.fail('Should have thrown an error');
        } catch (error) {
            assert.ok(error instanceof Error);
        }
    });

    await it('should get all services', async () => {
        await service.register('test', 'http://localhost:3000');
        await service.register('test2', 'http://localhost:3001');

        const services = await service.getServices();

        assert.strictEqual(services.length, 2);
    });

    await it('should delete a service', async () => {
        const { id } = await service.register('test', 'http://localhost:3000');
        await service.deleteService(id);

        const s = await service.getService('test');

        assert.strictEqual(s, undefined);
    });

    await it('should not be able to delete a service that does not exist', async () => {
        try {
            await service.deleteService('invalidId');
            assert.fail('Should have thrown an error');
        } catch (error) {
            assert.ok(error instanceof Error);
        }
    });
});
