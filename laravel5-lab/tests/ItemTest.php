<?php

use App\Item;
use Illuminate\Foundation\Testing\WithoutMiddleware;
use Illuminate\Foundation\Testing\DatabaseMigrations;

class ItemTest extends TestCase
{
    use DatabaseMigrations;

    public function testV1Index(){
        $this->seed('ItemTableSeeder');
        $response = $this->call('GET', '/v1/items');
        $this->assertThat($response->getStatusCode(), $this->equalTo(\Illuminate\Http\Response::HTTP_OK));
        $this->assertThat($response->getContent(), $this->stringContains('item-1'));
    }

    public function testV1IndexWithParam(){
        $this->seed('ItemTableSeeder');
        $response = $this->call('GET', '/v1/items?active=0');
        $this->assertThat($response->getStatusCode(), $this->equalTo(\Illuminate\Http\Response::HTTP_OK));
        $this->assertThat($response->getContent(), $this->equalTo('[]'));
        $response = $this->call('GET', '/v1/items?active=1');
        $this->assertThat($response->getStatusCode(), $this->equalTo(\Illuminate\Http\Response::HTTP_OK));
        $this->assertThat($response->getContent(), $this->stringContains('item-1'));

        $response = $this->call('GET', '/v1/items?active=1&name=test');
        $this->assertThat($response->getStatusCode(), $this->equalTo(\Illuminate\Http\Response::HTTP_OK));
        $this->assertThat($response->getContent(), $this->equalTo('[]'));
    }
}
