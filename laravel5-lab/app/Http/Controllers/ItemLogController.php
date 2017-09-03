<?php namespace App\Http\Controllers;

use App\Http\Requests;
use App\Http\Controllers\Controller;

use App\ItemLog;
use Illuminate\Http\Request;

class ItemLogController extends Controller {

	/**
	 * Display a listing of the resource.
	 *
	 * @return Response
	 */
	public function index()
	{
		$item_logs = ItemLog::orderBy('id', 'desc')->paginate(10);

		return view('item_logs.index', compact('item_logs'));
	}

	/**
	 * Show the form for creating a new resource.
	 *
	 * @return Response
	 */
	public function create()
	{
		return view('item_logs.create');
	}

	/**
	 * Store a newly created resource in storage.
	 *
	 * @param Request $request
	 * @return Response
	 */
	public function store(Request $request)
	{
		$item_log = new ItemLog();

		$item_log->action = $request->input("action");
        $item_log->item_id = $request->input("item_id");

		$item_log->save();

		return redirect()->route('item_logs.index')->with('message', 'Item created successfully.');
	}

	/**
	 * Display the specified resource.
	 *
	 * @param  int  $id
	 * @return Response
	 */
	public function show($id)
	{
		$item_log = ItemLog::findOrFail($id);

		return view('item_logs.show', compact('item_log'));
	}

	/**
	 * Show the form for editing the specified resource.
	 *
	 * @param  int  $id
	 * @return Response
	 */
	public function edit($id)
	{
		$item_log = ItemLog::findOrFail($id);

		return view('item_logs.edit', compact('item_log'));
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
		$item_log = ItemLog::findOrFail($id);

		$item_log->action = $request->input("action");
        $item_log->item_id = $request->input("item_id");

		$item_log->save();

		return redirect()->route('item_logs.index')->with('message', 'Item updated successfully.');
	}

	/**
	 * Remove the specified resource from storage.
	 *
	 * @param  int  $id
	 * @return Response
	 */
	public function destroy($id)
	{
		$item_log = ItemLog::findOrFail($id);
		$item_log->delete();

		return redirect()->route('item_logs.index')->with('message', 'Item deleted successfully.');
	}

}
