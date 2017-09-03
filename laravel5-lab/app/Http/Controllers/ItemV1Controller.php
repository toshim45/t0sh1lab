<?php namespace App\Http\Controllers;

use App\Http\Requests;
use App\Http\Controllers\Controller;

use App\Http\Services\ItemService;
use App\Item;
use Illuminate\Http\Request;
use Illuminate\Http\Response;

class ItemV1Controller extends Controller {

    private $itemService;

    function __construct(ItemService $itemService)
    {
        $this->itemService = $itemService;
    }

    /**
     * Display a listing of the resource.
     *
     * @param Request $request
     * @return Response
     */
	public function index(Request $request)
	{
        return $this->itemService->find($request)->take(10)->get();
	}

	/**
	 * Store a newly created resource in storage.
	 *
	 * @param Request $request
	 * @return Response
	 */
	public function store(Request $request)
	{
        $this->validate($request, Item::rules());
        $item = $this->itemService->store($request);

		return response()->json($item->getAttributes(), Response::HTTP_CREATED);
	}

	/**
	 * Display the specified resource.
	 *
	 * @param  int  $id
	 * @return Response
	 */
	public function show($id)
	{
		return $this->itemService->findOne($id);
	}

	/**
	 * Update the specified resource in storage.
	 *
	 * @param  int  $id
	 * @param Request $request
	 * @return Response
	 */
	public function update(Request $request, $id)
	{
        $this->validate($request, Item::rules());
        $item = $this->itemService->update($request, $id);
		return response()->json($item->getAttributes());
	}

	/**
	 * Remove the specified resource from storage.
	 *
	 * @param  int  $id
	 * @return Response
	 */
	public function destroy($id)
	{
        $this->itemService->delete($id);
		return response()->json('', Response::HTTP_NO_CONTENT);
	}

}
