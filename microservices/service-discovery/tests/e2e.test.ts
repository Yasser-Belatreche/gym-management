import supertest from 'supertest';
import assert from 'node:assert';
import { describe, it } from 'node:test';

const request = supertest('http://localhost:3000');

await describe('Service Discovery E2E Tests', async () => {
    let serviceId: string;
    const API_SECRET = process.env.API_SECRET!;

    await it('should register a new service', async () => {
        const response = await request
            .post('/api/v1/services')
            .send({ name: 'test', url: 'http://localhost:3000' })
            .set('x-api-secret', API_SECRET);

        assert.strictEqual(response.status, 201);
        assert.strictEqual(response.body.name, 'test');
        assert.strictEqual(response.body.url, 'http://localhost:3000');

        serviceId = response.body.id;
    });

    await it('should retrieve a registered service', async () => {
        const response = await request
            .get(`/api/v1/services/${serviceId}`)
            .set('x-api-secret', API_SECRET);

        assert.strictEqual(response.status, 200);
        assert.strictEqual(response.body.id, serviceId);
    });

    await it('should update a registered service', async () => {
        const response = await request
            .put(`/api/v1/services/${serviceId}`)
            .send({ name: 'updated', url: 'http://localhost:3001' })
            .set('x-api-secret', API_SECRET);

        assert.strictEqual(response.status, 200);
    });

    await it('should retrieve all registered services and the registered service should exist', async () => {
        const response = await request.get('/api/v1/services').set('x-api-secret', API_SECRET);

        assert.strictEqual(response.status, 200);
        assert.ok(response.body.some((service: any) => service.id === serviceId));
    });

    await it('should delete a registered service', async () => {
        const response = await request
            .delete(`/api/v1/services/${serviceId}`)
            .set('x-api-secret', API_SECRET);

        assert.strictEqual(response.status, 204);
    });

    await it('should retrieve all registered services and the deleted service should not exist', async () => {
        const response = await request.get('/api/v1/services').set('x-api-secret', API_SECRET);

        assert.strictEqual(response.status, 200);
        assert.ok(!response.body.some((service: any) => service.id === serviceId));
    });
});
