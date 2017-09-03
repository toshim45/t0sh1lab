<?php
/**
 * User: artikow
 * at 3/12/16 7:08 PM
 *
 */

namespace App\Http\Services;


use App\Item;
use Illuminate\Http\Exception\HttpResponseException;
use Illuminate\Http\JsonResponse;
use Illuminate\Http\Response;

interface Service {

    public function find($request);

    public function store($request);

    public function findOne($id);

    public function delete($id);

    public function update($request, $id);
}