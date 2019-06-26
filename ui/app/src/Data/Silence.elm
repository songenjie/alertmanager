{-
   Alertmanager API
   API of the Prometheus Alertmanager (https://github.com/prometheus/alertmanager)

   OpenAPI spec version: 0.0.1

   NOTE: This file is auto generated by the openapi-generator.
   https://github.com/openapitools/openapi-generator.git
   Do not edit this file manually.
-}


module Data.Silence exposing (Silence(..), decoder, encoder)

import Data.Matcher as Matcher exposing (Matcher)
import DateTime exposing (DateTime)
import Dict exposing (Dict)
import Json.Decode as Decode exposing (Decoder)
import Json.Decode.Pipeline exposing (optional, required)
import Json.Encode as Encode


type alias Silence =
    { matchers : List Matcher
    , startsAt : DateTime
    , endsAt : DateTime
    , createdBy : String
    , comment : String
    }


decoder : Decoder Silence
decoder =
    Decode.succeed Silence
        |> required "matchers" (Decode.list Matcher.decoder)
        |> required "startsAt" DateTime.decoder
        |> required "endsAt" DateTime.decoder
        |> required "createdBy" Decode.string
        |> required "comment" Decode.string


encoder : Silence -> Encode.Value
encoder model =
    Encode.object
        [ ( "matchers", Encode.list Matcher.encoder model.matchers )
        , ( "startsAt", DateTime.encoder model.startsAt )
        , ( "endsAt", DateTime.encoder model.endsAt )
        , ( "createdBy", Encode.string model.createdBy )
        , ( "comment", Encode.string model.comment )
        ]
