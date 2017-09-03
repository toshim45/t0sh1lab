<?php

namespace App;

use Illuminate\Database\Eloquent\Model;

class Item extends Model
{
    public static function rules()
    {
        return [
            'name' => 'required|max:255',
            'amount' => 'required',
            'active' => 'required',
        ];
    }
}
