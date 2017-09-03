<?php namespace App\Http\Controllers;

use App\Http\Requests;
use App\Http\Controllers\Controller;

use App\Http\Services\ItemService;
use App\Item;
use Illuminate\Http\Request;

class ItemController extends Controller {

    /* @var $itemService ItemService */
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
		$items = $this->itemService->find($request)->paginate(10);
		return view('items.index', compact('items'));
	}

    public function search(){
        return Item::all('id','name');
    }

	/**
	 * Show the form for creating a new resource.
	 *
	 * @return Response
	 */
	public function create()
	{
		return view('items.create');
	}

	/**
	 * Store a newly created resource in storage.
	 *
	 * @param Request $request
	 * @return Response
	 */
	public function store(Request $request)
	{
		$item = new Item();

		$item->name = $request->input("name");
        $item->amount = $request->input("amount");
        $item->active = $request->input("active");

		$item->save();

		return redirect()->route('items.index')->with('message', 'Item created successfully.');
	}

	/**
	 * Display the specified resource.
	 *
	 * @param  int  $id
	 * @return Response
	 */
	public function show($id)
	{
		$item = Item::findOrFail($id);

		return view('items.show', compact('item'));
	}

	/**
	 * Show the form for editing the specified resource.
	 *
	 * @param  int  $id
	 * @return Response
	 */
	public function edit($id)
	{
		$item = Item::findOrFail($id);

		return view('items.edit', compact('item'));
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
		$item = Item::findOrFail($id);

		$item->name = $request->input("name");
        $item->amount = $request->input("amount");
        $item->active = $request->input("active");

		$item->save();

		return redirect()->route('items.index')->with('message', 'Item updated successfully.');
	}

	/**
	 * Remove the specified resource from storage.
	 *
	 * @param  int  $id
	 * @return Response
	 */
	public function destroy($id)
	{
		$item = Item::findOrFail($id);
		$item->delete();

		return redirect()->route('items.index')->with('message', 'Item deleted successfully.');
	}

}
