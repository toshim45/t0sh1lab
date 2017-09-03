<?php namespace App\Http\Services;
use App\Item;
use Illuminate\Http\Exception\HttpResponseException;
use Illuminate\Http\JsonResponse;
use Illuminate\Http\Response;
use Symfony\Component\HttpFoundation\Request;

class ItemService implements Service
{
    /**
     * @param $request Request
     * @return mixed
     */
    public function find($request)
    {
        /* @var $item Item */
        $item = Item::query();

        $name = $request->query('name');
        $active = $request->query('active');
        if (isset($name)) $item->where('name','LIKE','%'.$name.'%');
        if (isset($active)) $item->where('active','=',$active);
        return $item;
    }

    public function store($request)
    {
        $item = new Item();
        return $this->save($request, $item);
    }

    /**
     * @param $request Request
     * @param $item Item
     * @return Item
     */
    public function save($request, $item)
    {
        $item->name = $request->input("name");
        $item->amount = $request->input("amount");
        $item->active = $request->input("active");
        if (!$item->save())
            throw new HttpResponseException(new JsonResponse(['persist FAIL'], Response::HTTP_INTERNAL_SERVER_ERROR));

        return $item;
    }

    public function findOne($id)
    {
        return Item::findOrfail($id);
    }

    public function delete($id)
    {
        /* @var $item Item */
        $item = Item::findOrFail($id);
        if (!$item->delete())
            throw new HttpResponseException(new JsonResponse(['delete FAIL'], Response::HTTP_INTERNAL_SERVER_ERROR));
    }

    public function update($request, $id)
    {
        $item = $this->findOne($id);
        return $this->save($request, $item);
    }
}
