{-
   Alertmanager API
   API of the Prometheus Alertmanager (https://github.com/prometheus/alertmanager)

   OpenAPI spec version: 0.0.1

   NOTE: This file is auto generated by the openapi-generator.
   https://github.com/openapitools/openapi-generator.git
   Do not edit this file manually.
-}


module Data.AlertGroup exposing (AlertGroup, decoder, encoder)

import Data.GettableAlert as GettableAlert exposing (GettableAlert)
import Data.Receiver as Receiver exposing (Receiver)
import Dict exposing (Dict)
import Json.Decode as Decode exposing (Decoder)
import Json.Decode.Pipeline exposing (optional, required)
import Json.Encode as Encode


type alias AlertGroup =
    { labels : Dict String String
    , receiver : Receiver
    , alerts : List GettableAlert
    }


decoder : Decoder AlertGroup
decoder =
    Decode.succeed AlertGroup
        |> required "labels" (Decode.dict Decode.string)
        |> required "receiver" Receiver.decoder
        |> required "alerts" (Decode.list GettableAlert.decoder)


encoder : AlertGroup -> Encode.Value
encoder model =
    Encode.object
        [ ( "labels", Encode.dict identity Encode.string model.labels )
        , ( "receiver", Receiver.encoder model.receiver )
        , ( "alerts", Encode.list GettableAlert.encoder model.alerts )
        ]
